package repository

type Aimprove struct {
	Id          int64     `db:"id" json:"id"`
	NamaLengkap string    `db:"nama_lengkap" json:"nama_lengkap"`
	Password    string    `db:"password" json:"password"`
	Email       string    `db:"email" json:"email"`
	PilihanCamp string    `db:"pilihan_camp" json:"pilihan_camp"`
	TanggalMulai string `db:"tanggal_mulai" json:"tanggal_mulai"`
	TanggalSelesai string `db:"tanggal_selesai" json:"tanggal_selesai"`
}

type User struct {
	Id          int64     `db:"id" json:"id"`
	Nama        string    `db:"nama" json:"nama"`
	Password    string    `db:"password" json:"password"`
	Email       string    `db:"email" json:"email"`
	NomorTelpon string `db:"no_telpon" json:"no_telpon"`
	TempatTanggalLahir string    `db:"tempat_tanggal_lahir" json:"tempat_tanggal_lahir"`
	Alamat         string    `db:"alamat" json:"alamat"`
	Pendidikan     string    `db:"pendidikan" json:"pendidikan"`
	
}

type Pendaftaran struct {
	Id          int64     `db:"id" json:"id"`
	IdAimprove  int64     `db:"id_aimprove" json:"id_aimprove"`
	IdUser      int64     `db:"id_user" json:"id_user"`
	TanggalDaftar string `db:"tanggal_daftar" json:"tanggal_daftar"`
	Status      string    `db:"status" json:"status"`
}

type Camp struct {
	Id          int64     `db:"id" json:"id"`
	NamaCamp    string    `db:"nama_camp" json:"nama_camp"`
	Email       string    `db:"email" json:"email"`
	Payment     string    `db:"payment" json:"payment"`
	Status      string    `db:"status" json:"status"`
}

type Iisma struct {
	Id          int64     `db:"id" json:"id"`
	PenjelasanIisma    string    `db:"penjelasan_iisma" json:"penjelasan_iisma"`
	BookletUnilist       string    `db:"booklet_unilist" json:"booklet_unilist"`
	InfoIisma    string    `db:"info_iisma" json:"info_iisma"`
}

type Career struct {
	Id          int64     `db:"id" json:"id"`
	PenjelasanCareer  string    `db:"penjelasan_career" json:"penjelasan_career"`
	BookletCareerlist       string    `db:"booklet_careerlist" json:"booklet_careerlist"`
	InfoCareer    string    `db:"info_career" json:"info_career"`
}

type Softskill struct {
	Id          int64     `db:"id" json:"id"`
	PenjelasanSoftskill    string    `db:"penjelasan_softskill" json:"penjelasan_softskill"`
	BookletSoftskill       string    `db:"booklet_softskill" json:"booklet_softskill"`
	InfoSoftskill    string    `db:"info_softskill" json:"info_softskill"`
}

