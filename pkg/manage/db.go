package manage

import (
	"time"
)

func (m *Manage) GetLastUpdateDaysAgo() (string, error) {

	var last time.Time

	rows, _ := m.DB.Model(Incident{}).Select("last_change_at as d").Order("1 desc").Limit(1).Rows()
	for rows.Next() {
		err := rows.Scan(&last)
		if err != nil {
			return "", err
		}
	}

	if last.IsZero() {
		return time.Now().AddDate(0, 0, -14).String(), nil
	}
	return last.String(), nil
}
