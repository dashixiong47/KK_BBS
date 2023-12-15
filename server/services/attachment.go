package services

func BuyAttachment(userId, attachmentId uint, number int) error {
	return AddIntegral(userId, 2, number, 6, attachmentId, "购买附件")
}
