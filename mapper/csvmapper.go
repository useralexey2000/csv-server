package mapper

import (
	"csv-server/domain"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mapRecord(v []string) (*domain.Record, error) {
	if len(v) < 49 {
		return nil, fmt.Errorf("malformed record")
	}
	rec := &domain.Record{}

	rec.Id, _ = strconv.Atoi(v[1])
	rec.Uid = v[2]
	rec.Domain = v[3]
	rec.Cn = v[4]
	rec.Department = v[5]
	rec.Title = v[6]
	rec.Who = v[7]
	rec.Logon_count, _ = strconv.Atoi(v[8])
	rec.Num_logons7, _ = strconv.Atoi(v[9])
	rec.Num_share7, _ = strconv.Atoi(v[10])
	rec.Num_file7, _ = strconv.Atoi(v[11])
	rec.Num_ad7, _ = strconv.Atoi(v[12])
	rec.Num_n7, _ = strconv.Atoi(v[13])
	rec.Num_logons14, _ = strconv.Atoi(v[14])
	rec.Num_share14, _ = strconv.Atoi(v[15])
	rec.Num_file14, _ = strconv.Atoi(v[16])
	rec.Num_ad14, _ = strconv.Atoi(v[17])
	rec.Num_n14, _ = strconv.Atoi(v[18])
	rec.Num_logons30, _ = strconv.Atoi(v[19])
	rec.Num_share30, _ = strconv.Atoi(v[20])
	rec.Num_file30, _ = strconv.Atoi(v[21])
	rec.Num_ad30, _ = strconv.Atoi(v[22])
	rec.Num_n30, _ = strconv.Atoi(v[23])
	rec.Num_logons150, _ = strconv.Atoi(v[24])
	rec.Num_share150, _ = strconv.Atoi(v[25])
	rec.Num_file150, _ = strconv.Atoi(v[26])
	rec.Num_ad150, _ = strconv.Atoi(v[27])
	rec.Num_n150, _ = strconv.Atoi(v[28])
	rec.Num_logons365, _ = strconv.Atoi(v[29])
	rec.Num_share365, _ = strconv.Atoi(v[30])
	rec.Num_file365, _ = strconv.Atoi(v[31])
	rec.Num_ad365, _ = strconv.Atoi(v[32])
	rec.Num_n365, _ = strconv.Atoi(v[33])
	rec.Has_user_principal_name, _ = strconv.Atoi(v[34])
	rec.Has_mail, _ = strconv.Atoi(v[35])
	rec.Has_phone, _ = strconv.Atoi(v[36])
	rec.Flag_disabled, _ = strconv.Atoi(v[37])
	rec.Flag_lockout, _ = strconv.Atoi(v[38])
	rec.Flag_password_not_required, _ = strconv.Atoi(v[39])
	rec.Flag_password_cant_change, _ = strconv.Atoi(v[40])
	rec.Flag_dont_expire_password, _ = strconv.Atoi(v[41])
	rec.Owned_files, _ = strconv.Atoi(v[42])
	rec.Num_mailboxes, _ = strconv.Atoi(v[43])
	rec.Num_member_of_groups, _ = strconv.Atoi(v[44])
	rec.Num_member_of_indirect_groups, _ = strconv.Atoi(v[45])
	rec.Member_of_indirect_groups_ids = groupsToSlice(v[46])
	rec.Member_of_groups_ids = groupsToSlice(v[47])
	rec.Is_admin, _ = strconv.Atoi(v[48])
	rec.Is_service, _ = strconv.Atoi(v[49])

	return rec, nil
}

func groupsToSlice(g string) []int {
	groups := strings.Split(g, ";")
	slGroups := make([]int, 0, len(groups))

	for _, v := range groups {
		g, _ := strconv.Atoi(v)
		slGroups = append(slGroups, g)
	}

	return slGroups
}

func Load(src string) ([]*domain.Record, error) {
	fr, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(fr)
	r.FieldsPerRecord = 50

	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(data) < 2 {
		return nil, fmt.Errorf("empty records")
	}

	recs := make([]*domain.Record, 0, len(data))
	for _, v := range data[1:] {
		rec, err := mapRecord(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		recs = append(recs, rec)
	}
	return recs, nil
}
