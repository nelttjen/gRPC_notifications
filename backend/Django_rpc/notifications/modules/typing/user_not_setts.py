import pydantic


class UserNotificationConfig(pydantic.BaseModel):
    user_id: int
    new_chapters: bool = pydantic.Field(default=True)
    special_offers: bool = pydantic.Field(default=True)
    comment_answer: bool = pydantic.Field(default=True)
    author_posts: bool = pydantic.Field(default=True)
    new_title_status: bool = pydantic.Field(default=True)
    new_achievements: bool = pydantic.Field(default=True)
    battlepass_new_level: bool = pydantic.Field(default=True)
    personal_recommendations: bool = pydantic.Field(default=True)
    new_messages: bool = pydantic.Field(default=True)
