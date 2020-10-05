package gorm

import (
	"fmt"
	"strings"

	"github.com/007lock/drath/contract"
	"github.com/jinzhu/gorm"
)

func criteriaApply(db *gorm.DB, crit *contract.RepoCriterias) *gorm.DB {
	tx := db
	for _, preload := range crit.Preloads {
		if preload.Criterias != nil {
			preloadCrit := preload.Criterias // Shadow variable
			tx = tx.Preload(preload.Relation, func(txr *gorm.DB) *gorm.DB {
				return criteriaApply(txr, preloadCrit)
			})
		} else {
			tx = tx.Preload(preload.Relation)
		}
	}
	if len(crit.Conditions) > 0 {
		tx = buildConditions(tx, crit.Conditions, false)
	}
	if len(crit.OrConditions) > 0 {
		tx = buildConditions(tx, crit.OrConditions, true)
	}

	for i := range crit.Joins {
		tx = tx.Joins(crit.Joins[i].Join)
	}

	for i := range crit.GroupBy {
		tx = tx.Group(crit.GroupBy[i])
	}

	for i := range crit.Order {
		tx = tx.Order(crit.Order[i])
	}
	return tx
}

func buildConditions(db *gorm.DB, conditions []*contract.RepoCondition, isOrCondition bool) *gorm.DB {
	tx := db
	parentOrEpr := ""
	for i, child := range conditions {
		if child.Field != "" {
			if isOrCondition {
				if i == 0 {
					parentOrEpr += fmt.Sprintf("%s %s", child.Field, child.Operation)
				} else {
					parentOrEpr += fmt.Sprintf(" OR %s %s", child.Field, child.Operation)
				}
				if child.Subquery {
					parentOrEpr += fmt.Sprintf(" (%s)", child.Value)
				} else if strings.ToLower(child.Operation) == "in" {
					parentOrEpr += " (?)"
				} else if child.Value != nil {
					parentOrEpr += " ?"
				} else {
					parentOrEpr += " NULL"
				}
			} else {
				if child.Subquery {
					tx = tx.Where(fmt.Sprintf("%s %s (%s)", child.Field, child.Operation, child.Value))
				} else if strings.ToLower(child.Operation) == "in" || strings.ToLower(child.Operation) == "not in" {
					tx = tx.Where(fmt.Sprintf("%s %s (?)", child.Field, child.Operation), child.Value)
				} else if child.Value != nil {
					tx = tx.Where(fmt.Sprintf("%s %s ?", child.Field, child.Operation), child.Value)
				} else {
					tx = tx.Where(fmt.Sprintf("%s %s NULL", child.Field, child.Operation))
				}
			}

		}
		for _, cond := range child.Conditions {
			if cond.Subquery {
				tx = tx.Where(fmt.Sprintf("%s %s (%s)", cond.Field, cond.Operation, cond.Value))
			} else if strings.ToLower(cond.Operation) == "in" {
				tx = tx.Where(fmt.Sprintf("%s %s (?)", cond.Field, cond.Operation), cond.Value)
			} else if cond.Value != nil {
				tx = tx.Where(fmt.Sprintf("%s %s ?", cond.Field, cond.Operation), cond.Value)
			} else {
				tx = tx.Where(fmt.Sprintf("%s %s NULL", cond.Field, cond.Operation))
			}
			if len(cond.Conditions) > 0 {
				tx = buildConditions(tx, cond.Conditions, false)
			}
			if len(cond.OrConditions) > 0 {
				tx = buildConditions(tx, cond.OrConditions, true)
			}
		}
		// Nested or
		orEpr := ""
		for i, cond := range child.OrConditions {
			if i == 0 {
				orEpr += fmt.Sprintf("%s %s", cond.Field, cond.Operation)
			} else {
				orEpr += fmt.Sprintf(" OR %s %s", cond.Field, cond.Operation)
			}
			if cond.Subquery {
				orEpr += fmt.Sprintf(" (%s)", cond.Value)
			} else if strings.ToLower(cond.Operation) == "in" {
				orEpr += " (?)"
			} else if cond.Value != nil {
				orEpr += " ?"
			} else {
				orEpr += " NULL"
			}
		}
		if orEpr != "" {
			var values []interface{}
			for _, cond := range child.OrConditions {
				if cond.Value != nil || !cond.Subquery {
					values = append(values, cond.Value)
				}
				if len(cond.Conditions) > 0 {
					tx = buildConditions(tx, cond.Conditions, false)
				}
				if len(cond.OrConditions) > 0 {
					tx = buildConditions(tx, cond.OrConditions, true)
				}
			}
			tx = tx.Where(orEpr, values...)
		}
	}

	if parentOrEpr != "" && isOrCondition {
		var values []interface{}
		for _, cond := range conditions {
			if cond.Value != nil || !cond.Subquery {
				values = append(values, cond.Value)
			}
			if len(cond.Conditions) > 0 {
				tx = buildConditions(tx, cond.Conditions, false)
			}
			if len(cond.OrConditions) > 0 {
				tx = buildConditions(tx, cond.OrConditions, false)
			}
		}
		tx = tx.Where(parentOrEpr, values...)
	}

	return tx
}
