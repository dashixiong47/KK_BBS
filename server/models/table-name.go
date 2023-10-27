package models

func (TopicBasic) TableName() string {
	return "kk_topic_basics"
}
func (TopicImage) TableName() string {
	return "kk_topic_images"
}
func (TopicText) TableName() string {
	return "kk_topic_texts"
}
func (TopicVideo) TableName() string {
	return "kk_topic_videos"
}
