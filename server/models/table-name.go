package models

func (TopicBasic) TableName() string {
	return "topic_basics"
}
func (TopicImage) TableName() string {
	return "topic_images"
}
func (TopicText) TableName() string {
	return "topic_texts"
}
func (TopicVideo) TableName() string {
	return "topic_videos"
}
