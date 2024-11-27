package unit

import (
	"testing"

	"example.com/sa-67-example/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestVaild(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`data valid`, func(t *testing.T) {
		member := entity.Member{
			Username: "testuser",
			Password: "123456",
			Email:    "test@example.com",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestUsername(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`username is required`, func(t *testing.T) {
		member := entity.Member{
			Username: "", // ผิดตรงนี้
			Password: "123456",
			Email:    "test@example.com",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Username is required"))
	})
}

func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`email format is invalid`, func(t *testing.T) {
		member := entity.Member{
			Username: "testuser",
			Password: "123456",
			Email:    "test-example.com", // รูปแบบผิด
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is invalid")) // ตรวจสอบข้อความ "Email is invalid"
	})
}