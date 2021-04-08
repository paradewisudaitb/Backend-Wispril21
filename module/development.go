package module

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Development(db *gorm.DB, g *gin.Engine) {
	g.GET("/dummy", middleware.Auth, func(c *gin.Context) {
		InsertDummy(db)
	})
	g.GET("/test/:id", func(c *gin.Context) {
		//results := []map[string]interface{}{}
		var results []entity.Wisudawan
		oid := c.Param("id")
		ids := db.Table("wisudawan").Joins("INNER JOIN content ON content.wisudawan_id = wisudawan.id").Joins("INNER JOIN organization ON organization.id = content.organization_id").Where("organization_id = ?", oid).Distinct("wisudawan.id")
		db.Preload(clause.Associations).Find(&results, "id IN (?)", ids)
		c.JSON(http.StatusOK, results)
	})
}

func InsertDummy(db *gorm.DB) {
	var JurusanDummy [3]entity.Jurusan
	JurusanDummy[0] = entity.Jurusan{
		Jurusan:       "Teknik Informatika",
		Fakultas:      "Sekolah Teknik Elektro Informatika",
		FakultasShort: "STEI",
		JurusanShort:  "HMIF",
	}
	JurusanDummy[1] = entity.Jurusan{
		Jurusan:       "Teknik Lingkungan",
		Fakultas:      "Fakultas Teknik Sipil dan Lingkungan",
		JurusanShort:  "TL",
		FakultasShort: "FTSL",
	}

	JurusanDummy[2] = entity.Jurusan{
		Jurusan:       "Matematika",
		Fakultas:      "Fakultas Matematika dan Ilmu Pengetahuan Alam",
		FakultasShort: "FMIPA",
		JurusanShort:  "MA",
	}
	db.Create(&JurusanDummy)

	var JurusanId [3]string
	JurusanId[0] = JurusanDummy[0].ID
	JurusanId[1] = JurusanDummy[1].ID
	JurusanId[2] = JurusanDummy[2].ID

	var OrganizationDummy [3]entity.Orgz
	OrganizationDummy[0] = entity.Orgz{
		Name:          "Himpunan Mahasiwswa Teknik Informatika",
		Slug:          "HMIF",
		Category:      "HMJ",
		Logo:          "/path/to/logo",
		FakultasShort: "STEI",
	}
	OrganizationDummy[1] = entity.Orgz{
		Name:     "Kabinet",
		Slug:     "Kabinet",
		Category: "Kabinet_KM_ITB",
		Logo:     "/path/to/logo",
	}
	OrganizationDummy[2] = entity.Orgz{
		Name:          "Himpunan Mahasiswa Elektro",
		Slug:          "HME",
		Category:      "HMJ",
		Logo:          "/path/to/logo",
		FakultasShort: "STEI",
	}

	db.Create(&OrganizationDummy)

	var OrganizationId [3]string
	OrganizationId[0] = OrganizationDummy[0].ID
	OrganizationId[1] = OrganizationDummy[1].ID
	OrganizationId[2] = OrganizationDummy[2].ID

	var WisudawanDummy [3]entity.Wisudawan
	WisudawanDummy[0] = entity.Wisudawan{
		Nim:          13519000,
		Nama:         "Sebuah nama",
		Panggilan:    "Sebuah panggilan",
		JudulTA:      "Belum TA",
		Angkatan:     16,
		JurusanID:    JurusanId[0],
		Instagram:    "gak ada",
		TanggalLahir: time.Now(),
		Photo:        "path/to/foto",
	}
	WisudawanDummy[1] = entity.Wisudawan{
		Nim:          13519001,
		Nama:         "Nama Lain",
		Panggilan:    "Panggilan Lain",
		JudulTA:      "Sudah TA",
		Angkatan:     15,
		JurusanID:    JurusanId[1],
		Instagram:    "ada",
		TanggalLahir: time.Now(),
		Photo:        "path/to/photo",
	}
	WisudawanDummy[2] = entity.Wisudawan{
		Nim:          13519002,
		Nama:         "Nama Lain lagi",
		Panggilan:    "Panggilan Lain lagi",
		JudulTA:      "Sudah TA lagi",
		Angkatan:     14,
		JurusanID:    JurusanId[2],
		Instagram:    "sudah ada",
		TanggalLahir: time.Now(),
		Photo:        "/path/to/newphoto",
	}
	db.Create(&WisudawanDummy)

	var idW [3]string
	idW[0] = WisudawanDummy[0].ID
	idW[1] = WisudawanDummy[1].ID
	idW[2] = WisudawanDummy[2].ID

	var MessageDummy [3]entity.Message
	MessageDummy[0] = entity.Message{
		ReceiverID: idW[0],
		Message:    "Halo",
		Sender:     "Anon",
	}

	MessageDummy[1] = entity.Message{
		ReceiverID: idW[0],
		Message:    "Miaw",
		Sender:     "Kucing ITB",
	}

	MessageDummy[2] = entity.Message{
		ReceiverID: idW[2],
		Message:    "Halo sayang",
		Sender:     "Secret Admirer",
	}
	db.Create(&MessageDummy)

	var ContentDummy [5]entity.Content

	ContentDummy[0] = entity.Content{
		WisudawanID: idW[0],
		Type:        "PRESTASI",
		Headings:    "Imba aku cuk",
	}
	ContentDummy[1] = entity.Content{
		WisudawanID: idW[0],
		Type:        "TIPS_SUKSES",
		Headings:    "Swimming aja",
		Details:     "Berenang menyehatkan badan",
	}
	ContentDummy[2] = entity.Content{
		WisudawanID:    idW[0],
		OrganizationID: OrganizationId[1],
		Type:           "KONTRIBUSI",
		Headings:       "Swimming aja",
	}
	ContentDummy[3] = entity.Content{
		WisudawanID:    idW[2],
		OrganizationID: OrganizationId[0],
		Type:           "KONTRIBUSI",
		Headings:       "Swimming aja",
	}
	ContentDummy[4] = entity.Content{
		WisudawanID:    idW[2],
		OrganizationID: OrganizationId[0],
		Type:           "KONTRIBUSI",
		Headings:       "Lalala",
	}
	db.Create(&ContentDummy)

}
