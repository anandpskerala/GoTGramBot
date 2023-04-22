package types

import "encoding/json"


// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
    // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
    UpdateId int64 `json:"update_id"`
    // Optional. New incoming message of any kind - text, photo, sticker, etc.
    Message *Message `json:"message,omitempty"`
    // Optional. New version of a message that is known to the bot and was edited
    EditedMessage *Message `json:"edited_message,omitempty"`
    // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
    ChannelPost *Message `json:"channel_post,omitempty"`
    // Optional. New version of a channel post that is known to the bot and was edited
    EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
    // Optional. New incoming inline query
    InlineQuery *InlineQuery `json:"inline_query,omitempty"`
    // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
    ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
    // Optional. New incoming callback query
    CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
    // Optional. New incoming shipping query. Only for invoices with flexible price
    ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
    // Optional. New incoming pre-checkout query. Contains full information about checkout
    PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
    // Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
    Poll *Poll `json:"poll,omitempty"`
    // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
    PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
    // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
    MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
    // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
    ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
    // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
    ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
}


// Describes the current status of a webhook.
type WebhookInfo struct {
    // Webhook URL, may be empty if webhook is not set up
    Url string `json:"url"`
    // True, if a custom certificate was provided for webhook certificate checks
    HasCustomCertificate bool `json:"has_custom_certificate"`
    // Number of updates awaiting delivery
    PendingUpdateCount int64 `json:"pending_update_count"`
    // Optional. Currently used webhook IP address
    IpAddress string `json:"ip_address,omitempty"`
    // Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
    LastErrorDate int64 `json:"last_error_date,omitempty"`
    // Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
    LastErrorMessage string `json:"last_error_message,omitempty"`
    // Optional. Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters
    LastSynchronizationErrorDate int64 `json:"last_synchronization_error_date,omitempty"`
    // Optional. The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
    MaxConnections int64 `json:"max_connections,omitempty"`
    // Optional. A list of update types the bot is subscribed to. Defaults to all update types except chat_member
    AllowedUpdates []string `json:"allowed_updates,omitempty"`
}


// This object represents a Telegram user or bot.
type User struct {
    // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
    Id int64 `json:"id"`
    // True, if this user is a bot
    IsBot bool `json:"is_bot"`
    // User's or bot's first name
    FirstName string `json:"first_name"`
    // Optional. User's or bot's last name
    LastName string `json:"last_name,omitempty"`
    // Optional. User's or bot's username
    Username string `json:"username,omitempty"`
    // Optional. IETF language tag of the user's language
    LanguageCode string `json:"language_code,omitempty"`
    // Optional. True, if this user is a Telegram Premium user
    IsPremium bool `json:"is_premium,omitempty"`
    // Optional. True, if this user added the bot to the attachment menu
    AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`
    // Optional. True, if the bot can be invited to groups. Returned only in getMe.
    CanJoinGroups bool `json:"can_join_groups,omitempty"`
    // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
    CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
    // Optional. True, if the bot supports inline queries. Returned only in getMe.
    SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}


// This object represents a chat.
type Chat struct {
    // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
    Id int64 `json:"id"`
    // Type of chat, can be either "private", "group", "supergroup" or "channel"
    Type string `json:"type"`
    // Optional. Title, for supergroups, channels and group chats
    Title string `json:"title,omitempty"`
    // Optional. Username, for private chats, supergroups and channels if available
    Username string `json:"username,omitempty"`
    // Optional. First name of the other party in a private chat
    FirstName string `json:"first_name,omitempty"`
    // Optional. Last name of the other party in a private chat
    LastName string `json:"last_name,omitempty"`
    // Optional. True, if the supergroup chat is a forum (has topics enabled)
    IsForum bool `json:"is_forum,omitempty"`
    // Optional. Chat photo. Returned only in getChat.
    Photo *ChatPhoto `json:"photo,omitempty"`
    // Optional. If non-empty, the list of all active chat usernames; for private chats, supergroups and channels. Returned only in getChat.
    ActiveUsernames []string `json:"active_usernames,omitempty"`
    // Optional. Custom emoji identifier of emoji status of the other party in a private chat. Returned only in getChat.
    EmojiStatusCustomEmojiId string `json:"emoji_status_custom_emoji_id,omitempty"`
    // Optional. Bio of the other party in a private chat. Returned only in getChat.
    Bio string `json:"bio,omitempty"`
    // Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
    HasPrivateForwards bool `json:"has_private_forwards,omitempty"`
    // Optional. True, if the privacy settings of the other party restrict sending voice and video note messages in the private chat. Returned only in getChat.
    HasRestrictedVoiceAndVideoMessages bool `json:"has_restricted_voice_and_video_messages,omitempty"`
    // Optional. True, if users need to join the supergroup before they can send messages. Returned only in getChat.
    JoinToSendMessages bool `json:"join_to_send_messages,omitempty"`
    // Optional. True, if all users directly joining the supergroup need to be approved by supergroup administrators. Returned only in getChat.
    JoinByRequest bool `json:"join_by_request,omitempty"`
    // Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
    Description string `json:"description,omitempty"`
    // Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
    InviteLink string `json:"invite_link,omitempty"`
    // Optional. The most recent pinned message (by sending date). Returned only in getChat.
    PinnedMessage *Message `json:"pinned_message,omitempty"`
    // Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
    Permissions *ChatPermissions `json:"permissions,omitempty"`
    // Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
    SlowModeDelay int64 `json:"slow_mode_delay,omitempty"`
    // Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
    MessageAutoDeleteTime int64 `json:"message_auto_delete_time,omitempty"`
    // Optional. True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available to chat administrators. Returned only in getChat.
    HasAggressiveAntiSpamEnabled bool `json:"has_aggressive_anti_spam_enabled,omitempty"`
    // Optional. True, if non-administrators can only get the list of bots and administrators in the chat. Returned only in getChat.
    HasHiddenMembers bool `json:"has_hidden_members,omitempty"`
    // Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
    HasProtectedContent bool `json:"has_protected_content,omitempty"`
    // Optional. For supergroups, name of group sticker set. Returned only in getChat.
    StickerSetName string `json:"sticker_set_name,omitempty"`
    // Optional. True, if the bot can change the group sticker set. Returned only in getChat.
    CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
    // Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
    LinkedChatId int64 `json:"linked_chat_id,omitempty"`
    // Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
    Location *ChatLocation `json:"location,omitempty"`
}


// This object represents a message.
type Message struct {
    // Unique message identifier inside this chat
    MessageId int64 `json:"message_id"`
    // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
    MessageThreadId int64 `json:"message_thread_id,omitempty"`
    // Optional. Sender of the message; empty for messages sent to channels. For backward compatibility, the field contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
    From *User `json:"from,omitempty"`
    // Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for channel posts, the supergroup itself for messages from anonymous group administrators, the linked channel for messages automatically forwarded to the discussion group. For backward compatibility, the field from contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
    SenderChat *Chat `json:"sender_chat,omitempty"`
    // Date the message was sent in Unix time
    Date int64 `json:"date"`
    // Conversation the message belongs to
    Chat *Chat `json:"chat"`
    // Optional. For forwarded messages, sender of the original message
    ForwardFrom *User `json:"forward_from,omitempty"`
    // Optional. For messages forwarded from channels or from anonymous administrators, information about the original sender chat
    ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
    // Optional. For messages forwarded from channels, identifier of the original message in the channel
    ForwardFromMessageId int64 `json:"forward_from_message_id,omitempty"`
    // Optional. For forwarded messages that were originally sent in channels or by an anonymous chat administrator, signature of the message sender if present
    ForwardSignature string `json:"forward_signature,omitempty"`
    // Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
    ForwardSenderName string `json:"forward_sender_name,omitempty"`
    // Optional. For forwarded messages, date the original message was sent in Unix time
    ForwardDate int64 `json:"forward_date,omitempty"`
    // Optional. True, if the message is sent to a forum topic
    IsTopicMessage bool `json:"is_topic_message,omitempty"`
    // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
    IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`
    // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
    ReplyToMessage *Message `json:"reply_to_message,omitempty"`
    // Optional. Bot through which the message was sent
    ViaBot *User `json:"via_bot,omitempty"`
    // Optional. Date the message was last edited in Unix time
    EditDate int64 `json:"edit_date,omitempty"`
    // Optional. True, if the message can't be forwarded
    HasProtectedContent bool `json:"has_protected_content,omitempty"`
    // Optional. The unique identifier of a media message group this message belongs to
    MediaGroupId string `json:"media_group_id,omitempty"`
    // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
    AuthorSignature string `json:"author_signature,omitempty"`
    // Optional. For text messages, the actual UTF-8 text of the message
    Text string `json:"text,omitempty"`
    // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
    Entities []MessageEntity `json:"entities,omitempty"`
    // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
    Animation *Animation `json:"animation,omitempty"`
    // Optional. Message is an audio file, information about the file
    Audio *Audio `json:"audio,omitempty"`
    // Optional. Message is a general file, information about the file
    Document *Document `json:"document,omitempty"`
    // Optional. Message is a photo, available sizes of the photo
    Photo []PhotoSize `json:"photo,omitempty"`
    // Optional. Message is a sticker, information about the sticker
    Sticker *Sticker `json:"sticker,omitempty"`
    // Optional. Message is a video, information about the video
    Video *Video `json:"video,omitempty"`
    // Optional. Message is a video note, information about the video message
    VideoNote *VideoNote `json:"video_note,omitempty"`
    // Optional. Message is a voice message, information about the file
    Voice *Voice `json:"voice,omitempty"`
    // Optional. Caption for the animation, audio, document, photo, video or voice
    Caption string `json:"caption,omitempty"`
    // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. True, if the message media is covered by a spoiler animation
    HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`
    // Optional. Message is a shared contact, information about the contact
    Contact *Contact `json:"contact,omitempty"`
    // Optional. Message is a dice with random value
    Dice *Dice `json:"dice,omitempty"`
    // Optional. Message is a game, information about the game. More about games: https://core.telegram.org/bots/api#games
    Game *Game `json:"game,omitempty"`
    // Optional. Message is a native poll, information about the poll
    Poll *Poll `json:"poll,omitempty"`
    // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
    Venue *Venue `json:"venue,omitempty"`
    // Optional. Message is a shared location, information about the location
    Location *Location `json:"location,omitempty"`
    // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
    NewChatMembers []User `json:"new_chat_members,omitempty"`
    // Optional. A member was removed from the group, information about them (this member may be the bot itself)
    LeftChatMember *User `json:"left_chat_member,omitempty"`
    // Optional. A chat title was changed to this value
    NewChatTitle string `json:"new_chat_title,omitempty"`
    // Optional. A chat photo was change to this value
    NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`
    // Optional. Service message: the chat photo was deleted
    DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
    // Optional. Service message: the group has been created
    GroupChatCreated bool `json:"group_chat_created,omitempty"`
    // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
    SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
    // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
    ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
    // Optional. Service message: auto-delete timer settings changed in the chat
    MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
    // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
    MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
    // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
    MigrateFromChatId int64 `json:"migrate_from_chat_id,omitempty"`
    // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
    PinnedMessage *Message `json:"pinned_message,omitempty"`
    // Optional. Message is an invoice for a payment, information about the invoice. More about payments: https://core.telegram.org/bots/api#payments
    Invoice *Invoice `json:"invoice,omitempty"`
    // Optional. Message is a service message about a successful payment, information about the payment. More about payments: https://core.telegram.org/bots/api#payments
    SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
    // Optional. Service message: a user was shared with the bot
    UserShared *UserShared `json:"user_shared,omitempty"`
    // Optional. Service message: a chat was shared with the bot
    ChatShared *ChatShared `json:"chat_shared,omitempty"`
    // Optional. The domain name of the website on which the user has logged in. More about Telegram Login: https://core.telegram.org/widgets/login
    ConnectedWebsite string `json:"connected_website,omitempty"`
    // Optional. Service message: the user allowed the bot added to the attachment menu to write messages
    WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`
    // Optional. Telegram Passport data
    PassportData *PassportData `json:"passport_data,omitempty"`
    // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
    ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
    // Optional. Service message: forum topic created
    ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`
    // Optional. Service message: forum topic edited
    ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`
    // Optional. Service message: forum topic closed
    ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`
    // Optional. Service message: forum topic reopened
    ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`
    // Optional. Service message: the 'General' forum topic hidden
    GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`
    // Optional. Service message: the 'General' forum topic unhidden
    GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`
    // Optional. Service message: video chat scheduled
    VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
    // Optional. Service message: video chat started
    VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`
    // Optional. Service message: video chat ended
    VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
    // Optional. Service message: new participants invited to a video chat
    VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
    // Optional. Service message: data sent by a Web App
    WebAppData *WebAppData `json:"web_app_data,omitempty"`
    // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}


// This object represents a unique message identifier.
type MessageId struct {
    // Unique message identifier
    MessageId int64 `json:"message_id"`
}


// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
    // Type of the entity. Currently, can be "mention" (@username), "hashtag" (#hashtag), "cashtag" ($USD), "bot_command" (/start@jobs_bot), "url" (https://telegram.org), "email" (do-not-reply@telegram.org), "phone_number" (+1-212-555-0123), "bold" (bold text), "italic" (italic text), "underline" (underlined text), "strikethrough" (strikethrough text), "spoiler" (spoiler message), "code" (monowidth string), "pre" (monowidth block), "text_link" (for clickable text URLs), "text_mention" (for users without usernames), "custom_emoji" (for inline custom emoji stickers)
    Type string `json:"type"`
    // Offset in UTF-16 code units to the start of the entity
    Offset int64 `json:"offset"`
    // Length of the entity in UTF-16 code units
    Length int64 `json:"length"`
    // Optional. For "text_link" only, URL that will be opened after user taps on the text
    Url string `json:"url,omitempty"`
    // Optional. For "text_mention" only, the mentioned user
    User *User `json:"user,omitempty"`
    // Optional. For "pre" only, the programming language of the entity text
    Language string `json:"language,omitempty"`
    // Optional. For "custom_emoji" only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
    CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}


// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Photo width
    Width int64 `json:"width"`
    // Photo height
    Height int64 `json:"height"`
    // Optional. File size in bytes
    FileSize int64 `json:"file_size,omitempty"`
}


// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Video width as defined by sender
    Width int64 `json:"width"`
    // Video height as defined by sender
    Height int64 `json:"height"`
    // Duration of the video in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Optional. Animation thumbnail as defined by sender
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
    // Optional. Original animation filename as defined by sender
    FileName string `json:"file_name,omitempty"`
    // Optional. MIME type of the file as defined by sender
    MimeType string `json:"mime_type,omitempty"`
    // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
    FileSize int64 `json:"file_size,omitempty"`
}


// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Duration of the audio in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Optional. Performer of the audio as defined by sender or by audio tags
    Performer string `json:"performer,omitempty"`
    // Optional. Title of the audio as defined by sender or by audio tags
    Title string `json:"title,omitempty"`
    // Optional. Original filename as defined by sender
    FileName string `json:"file_name,omitempty"`
    // Optional. MIME type of the file as defined by sender
    MimeType string `json:"mime_type,omitempty"`
    // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
    FileSize int64 `json:"file_size,omitempty"`
    // Optional. Thumbnail of the album cover to which the music file belongs
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}


// This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Optional. Document thumbnail as defined by sender
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
    // Optional. Original filename as defined by sender
    FileName string `json:"file_name,omitempty"`
    // Optional. MIME type of the file as defined by sender
    MimeType string `json:"mime_type,omitempty"`
    // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
    FileSize int64 `json:"file_size,omitempty"`
}


// This object represents a video file.
type Video struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Video width as defined by sender
    Width int64 `json:"width"`
    // Video height as defined by sender
    Height int64 `json:"height"`
    // Duration of the video in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Optional. Video thumbnail
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
    // Optional. Original filename as defined by sender
    FileName string `json:"file_name,omitempty"`
    // Optional. MIME type of the file as defined by sender
    MimeType string `json:"mime_type,omitempty"`
    // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
    FileSize int64 `json:"file_size,omitempty"`
}


// This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Video width and height (diameter of the video message) as defined by sender
    Length int64 `json:"length"`
    // Duration of the video in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Optional. Video thumbnail
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
    // Optional. File size in bytes
    FileSize int64 `json:"file_size,omitempty"`
}


// This object represents a voice note.
type Voice struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Duration of the audio in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Optional. MIME type of the file as defined by sender
    MimeType string `json:"mime_type,omitempty"`
    // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
    FileSize int64 `json:"file_size,omitempty"`
}


// This object represents a phone contact.
type Contact struct {
    // Contact's phone number
    PhoneNumber string `json:"phone_number"`
    // Contact's first name
    FirstName string `json:"first_name"`
    // Optional. Contact's last name
    LastName string `json:"last_name,omitempty"`
    // Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
    UserId int64 `json:"user_id,omitempty"`
    // Optional. Additional data about the contact in the form of a vCard
    Vcard string `json:"vcard,omitempty"`
}


// This object represents an animated emoji that displays a random value.
type Dice struct {
    // Emoji on which the dice throw animation is based
    Emoji string `json:"emoji"`
    // Value of the dice, 1-6 for "🎲", "🎯" and "🎳" base emoji, 1-5 for "🏀" and "⚽" base emoji, 1-64 for "🎰" base emoji
    Value int64 `json:"value"`
}


// This object contains information about one answer option in a poll.
type PollOption struct {
    // Option text, 1-100 characters
    Text string `json:"text"`
    // Number of users that voted for this option
    VoterCount int64 `json:"voter_count"`
}


// This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
    // Unique poll identifier
    PollId string `json:"poll_id"`
    // The user, who changed the answer to the poll
    User *User `json:"user"`
    // 0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.
    OptionIds []int64 `json:"option_ids"`
}


// This object contains information about a poll.
type Poll struct {
    // Unique poll identifier
    Id string `json:"id"`
    // Poll question, 1-300 characters
    Question string `json:"question"`
    // List of poll options
    Options []PollOption `json:"options"`
    // Total number of users that voted in the poll
    TotalVoterCount int64 `json:"total_voter_count"`
    // True, if the poll is closed
    IsClosed bool `json:"is_closed"`
    // True, if the poll is anonymous
    IsAnonymous bool `json:"is_anonymous"`
    // Poll type, currently can be "regular" or "quiz"
    Type string `json:"type"`
    // True, if the poll allows multiple answers
    AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
    // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
    CorrectOptionId int64 `json:"correct_option_id,omitempty"`
    // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
    Explanation string `json:"explanation,omitempty"`
    // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
    ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
    // Optional. Amount of time in seconds the poll will be active after creation
    OpenPeriod int64 `json:"open_period,omitempty"`
    // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
    CloseDate int64 `json:"close_date,omitempty"`
}


// This object represents a point on the map.
type Location struct {
    // Longitude as defined by sender
    Longitude float64 `json:"longitude"`
    // Latitude as defined by sender
    Latitude float64 `json:"latitude"`
    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    // Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
    LivePeriod int64 `json:"live_period,omitempty"`
    // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
    Heading int64 `json:"heading,omitempty"`
    // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
    ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
}


// This object represents a venue.
type Venue struct {
    // Venue location. Can't be a live location
    Location *Location `json:"location"`
    // Name of the venue
    Title string `json:"title"`
    // Address of the venue
    Address string `json:"address"`
    // Optional. Foursquare identifier of the venue
    FoursquareId string `json:"foursquare_id,omitempty"`
    // Optional. Foursquare type of the venue. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
    FoursquareType string `json:"foursquare_type,omitempty"`
    // Optional. Google Places identifier of the venue
    GooglePlaceId string `json:"google_place_id,omitempty"`
    // Optional. Google Places type of the venue. (See supported types.)
    GooglePlaceType string `json:"google_place_type,omitempty"`
}


// Describes data sent from a Web App to the bot.
type WebAppData struct {
    // The data. Be aware that a bad client can send arbitrary data in this field.
    Data string `json:"data"`
    // Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
    ButtonText string `json:"button_text"`
}


// This object represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
    // User that triggered the alert
    Traveler *User `json:"traveler"`
    // User that set the alert
    Watcher *User `json:"watcher"`
    // The distance between the users
    Distance int64 `json:"distance"`
}


// This object represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
    // New auto-delete time for messages in the chat; in seconds
    MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`
}


// This object represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
    // Name of the topic
    Name string `json:"name"`
    // Color of the topic icon in RGB format
    IconColor int64 `json:"icon_color"`
    // Optional. Unique identifier of the custom emoji shown as the topic icon
    IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}


// This object represents a service message about a forum topic closed in the chat. Currently holds no information.
type ForumTopicClosed interface {

}


// This object represents a service message about an edited forum topic.
type ForumTopicEdited struct {
    // Optional. New name of the topic, if it was edited
    Name string `json:"name,omitempty"`
    // Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
    IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}


// This object represents a service message about a forum topic reopened in the chat. Currently holds no information.
type ForumTopicReopened interface {

}


// This object represents a service message about General forum topic hidden in the chat. Currently holds no information.
type GeneralForumTopicHidden interface {

}


// This object represents a service message about General forum topic unhidden in the chat. Currently holds no information.
type GeneralForumTopicUnhidden interface {

}


// This object contains information about the user whose identifier was shared with the bot using a KeyboardButtonRequestUser button.
type UserShared struct {
    // Identifier of the request
    RequestId int64 `json:"request_id"`
    // Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
    UserId int64 `json:"user_id"`
}


// This object contains information about the chat whose identifier was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
    // Identifier of the request
    RequestId int64 `json:"request_id"`
    // Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
    ChatId int64 `json:"chat_id"`
}


// This object represents a service message about a user allowing a bot to write messages after adding the bot to the attachment menu or launching a Web App from a link.
type WriteAccessAllowed struct {
    // Optional. Name of the Web App which was launched from a link
    WebAppName string `json:"web_app_name,omitempty"`
}


// This object represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
    // Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
    StartDate int64 `json:"start_date"`
}


// This object represents a service message about a video chat started in the chat. Currently holds no information.
type VideoChatStarted interface {

}


// This object represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
    // Video chat duration in seconds
    Duration int64 `json:"duration"`
}


// This object represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
    // New members that were invited to the video chat
    Users []User `json:"users"`
}


// This object represent a user's profile pictures.
type UserProfilePhotos struct {
    // Total number of profile pictures the target user has
    TotalCount int64 `json:"total_count"`
    // Requested profile pictures (in up to 4 sizes each)
    Photos [][]PhotoSize `json:"photos"`
}


// This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
type File struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
    FileSize int64 `json:"file_size,omitempty"`
    // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
    FilePath string `json:"file_path,omitempty"`
}


// Describes a Web App.
type WebAppInfo struct {
    // An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
    Url string `json:"url"`
}


// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
    // Array of button rows, each represented by an Array of KeyboardButton objects
    Keyboard [][]KeyboardButton `json:"keyboard"`
    // Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
    IsPersistent bool `json:"is_persistent,omitempty"`
    // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
    ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
    // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
    OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
    // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
    InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
    // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
    Selective bool `json:"selective,omitempty"`
}


func (m ReplyKeyboardMarkup) replyMarkup() {}

// This object represents one button of the reply keyboard. For simple text buttons, String can be used instead of this object to specify the button text. The optional fields web_app, request_user, request_chat, request_contact, request_location, and request_poll are mutually exclusive.
// Note: request_contact and request_location options will only work in Telegram versions released after 9 April, 2016. Older clients will display unsupported message.
// Note: request_poll option will only work in Telegram versions released after 23 January, 2020. Older clients will display unsupported message.
// Note: web_app option will only work in Telegram versions released after 16 April, 2022. Older clients will display unsupported message.
// Note: request_user and request_chat options will only work in Telegram versions released after 3 February, 2023. Older clients will display unsupported message.
type KeyboardButton struct {
    // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
    Text string `json:"text"`
    // Optional. If specified, pressing the button will open a list of suitable users. Tapping on any user will send their identifier to the bot in a "user_shared" service message. Available in private chats only.
    RequestUser *KeyboardButtonRequestUser `json:"request_user,omitempty"`
    // Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a "chat_shared" service message. Available in private chats only.
    RequestChat *KeyboardButtonRequestChat `json:"request_chat,omitempty"`
    // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
    RequestContact bool `json:"request_contact,omitempty"`
    // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
    RequestLocation bool `json:"request_location,omitempty"`
    // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
    RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
    // Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a "web_app_data" service message. Available in private chats only.
    WebApp *WebAppInfo `json:"web_app,omitempty"`
}


// This object defines the criteria used to request a suitable user. The identifier of the selected user will be shared with the bot when the corresponding button is pressed. More about requesting users: https://core.telegram.org/bots/features#chat-and-user-selection
type KeyboardButtonRequestUser struct {
    // Signed 32-bit identifier of the request, which will be received back in the UserShared object. Must be unique within the message
    RequestId int64 `json:"request_id"`
    // Optional. Pass True to request a bot, pass False to request a regular user. If not specified, no additional restrictions are applied.
    UserIsBot bool `json:"user_is_bot,omitempty"`
    // Optional. Pass True to request a premium user, pass False to request a non-premium user. If not specified, no additional restrictions are applied.
    UserIsPremium bool `json:"user_is_premium,omitempty"`
}


// This object defines the criteria used to request a suitable chat. The identifier of the selected chat will be shared with the bot when the corresponding button is pressed. More about requesting chats: https://core.telegram.org/bots/features#chat-and-user-selection
type KeyboardButtonRequestChat struct {
    // Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
    RequestId int64 `json:"request_id"`
    // Pass True to request a channel chat, pass False to request a group or a supergroup chat.
    ChatIsChannel bool `json:"chat_is_channel"`
    // Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
    ChatIsForum bool `json:"chat_is_forum,omitempty"`
    // Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
    ChatHasUsername bool `json:"chat_has_username,omitempty"`
    // Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
    ChatIsCreated bool `json:"chat_is_created,omitempty"`
    // Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
    UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
    // Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
    BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
    // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
    BotIsMember bool `json:"bot_is_member,omitempty"`
}


// This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
    // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
    Type string `json:"type,omitempty"`
}


// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
    // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
    RemoveKeyboard bool `json:"remove_keyboard"`
    // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
    Selective bool `json:"selective,omitempty"`
}


func (m ReplyKeyboardRemove) replyMarkup() {}

// This object represents an inline keyboard that appears right next to the message it belongs to.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will display unsupported message.
type InlineKeyboardMarkup struct {
    // Array of button rows, each represented by an Array of InlineKeyboardButton objects
    InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}


func (m InlineKeyboardMarkup) replyMarkup() {}

// This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
    // Label text on the button
    Text string `json:"text"`
    // Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their ID without using a username, if this is allowed by their privacy settings.
    Url string `json:"url,omitempty"`
    // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
    CallbackData string `json:"callback_data,omitempty"`
    // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot.
    WebApp *WebAppInfo `json:"web_app,omitempty"`
    // Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
    LoginUrl *LoginUrl `json:"login_url,omitempty"`
    // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted. Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm... actions - in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
    SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
    // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted. This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options.
    SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
    // Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field
    SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
    // Optional. Description of the game that will be launched when the user presses the button. NOTE: This type of button must always be the first button in the first row.
    CallbackGame *CallbackGame `json:"callback_game,omitempty"`
    // Optional. Specify True, to send a Pay button. NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
    Pay bool `json:"pay,omitempty"`
}


// This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:
// Telegram apps support these buttons as of version 5.7.
type LoginUrl struct {
    // An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data. NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
    Url string `json:"url"`
    // Optional. New text of the button in forwarded messages.
    ForwardText string `json:"forward_text,omitempty"`
    // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
    BotUsername string `json:"bot_username,omitempty"`
    // Optional. Pass True to request the permission for your bot to send messages to the user.
    RequestWriteAccess bool `json:"request_write_access,omitempty"`
}


// This object represents an inline button that switches the current user to inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
    // Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
    Query string `json:"query,omitempty"`
    // Optional. True, if private chats with users can be chosen
    AllowUserChats bool `json:"allow_user_chats,omitempty"`
    // Optional. True, if private chats with bots can be chosen
    AllowBotChats bool `json:"allow_bot_chats,omitempty"`
    // Optional. True, if group and supergroup chats can be chosen
    AllowGroupChats bool `json:"allow_group_chats,omitempty"`
    // Optional. True, if channel chats can be chosen
    AllowChannelChats bool `json:"allow_channel_chats,omitempty"`
}


// This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
    // Unique identifier for this query
    Id string `json:"id"`
    // Sender
    From *User `json:"from"`
    // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
    Message *Message `json:"message,omitempty"`
    // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
    ChatInstance string `json:"chat_instance"`
    // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
    Data string `json:"data,omitempty"`
    // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
    GameShortName string `json:"game_short_name,omitempty"`
}


// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
    // Shows reply interface to the user, as if they manually selected the bot's message and tapped 'Reply'
    ForceReply bool `json:"force_reply"`
    // Optional. The placeholder to be shown in the input field when the reply is active; 1-64 characters
    InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
    // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
    Selective bool `json:"selective,omitempty"`
}


func (m ForceReply) replyMarkup() {}

// This object represents a chat photo.
type ChatPhoto struct {
    // File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
    SmallFileId string `json:"small_file_id"`
    // Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    SmallFileUniqueId string `json:"small_file_unique_id"`
    // File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
    BigFileId string `json:"big_file_id"`
    // Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    BigFileUniqueId string `json:"big_file_unique_id"`
}


// Represents an invite link for a chat.
type ChatInviteLink struct {
    // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with "...".
    InviteLink string `json:"invite_link"`
    // Creator of the link
    Creator *User `json:"creator"`
    // True, if users joining the chat via the link need to be approved by chat administrators
    CreatesJoinRequest bool `json:"creates_join_request"`
    // True, if the link is primary
    IsPrimary bool `json:"is_primary"`
    // True, if the link is revoked
    IsRevoked bool `json:"is_revoked"`
    // Optional. Invite link name
    Name string `json:"name,omitempty"`
    // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
    ExpireDate int64 `json:"expire_date,omitempty"`
    // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
    MemberLimit int64 `json:"member_limit,omitempty"`
    // Optional. Number of pending join requests created using this link
    PendingJoinRequestCount int64 `json:"pending_join_request_count,omitempty"`
}


// Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
    // True, if the user's presence in the chat is hidden
    IsAnonymous bool `json:"is_anonymous"`
    // True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
    CanManageChat bool `json:"can_manage_chat"`
    // True, if the administrator can delete messages of other users
    CanDeleteMessages bool `json:"can_delete_messages"`
    // True, if the administrator can manage video chats
    CanManageVideoChats bool `json:"can_manage_video_chats"`
    // True, if the administrator can restrict, ban or unban chat members
    CanRestrictMembers bool `json:"can_restrict_members"`
    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
    CanPromoteMembers bool `json:"can_promote_members"`
    // True, if the user is allowed to change the chat title, photo and other settings
    CanChangeInfo bool `json:"can_change_info"`
    // True, if the user is allowed to invite new users to the chat
    CanInviteUsers bool `json:"can_invite_users"`
    // Optional. True, if the administrator can post in the channel; channels only
    CanPostMessages bool `json:"can_post_messages,omitempty"`
    // Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
    CanEditMessages bool `json:"can_edit_messages,omitempty"`
    // Optional. True, if the user is allowed to pin messages; groups and supergroups only
    CanPinMessages bool `json:"can_pin_messages,omitempty"`
    // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
    CanManageTopics bool `json:"can_manage_topics,omitempty"`
}


// This object contains information about one member of a chat. Currently, the following 6 types of chat members are supported:
// - ChatMemberOwner
// - ChatMemberAdministrator
// - ChatMemberMember
// - ChatMemberRestricted
// - ChatMemberLeft
// - ChatMemberBanned
type ChatMember struct {
    Status string `json:"status"`
    User *User `json:"user"`
    IsAnonymous bool `json:"is_anonymous"`
    CustomTitle string `json:"custom_title,omitempty"`
    CanBeEdited bool `json:"can_be_edited"`
    CanManageChat bool `json:"can_manage_chat"`
    CanDeleteMessages bool `json:"can_delete_messages"`
    CanManageVideoChats bool `json:"can_manage_video_chats"`
    CanRestrictMembers bool `json:"can_restrict_members"`
    CanPromoteMembers bool `json:"can_promote_members"`
    CanChangeInfo bool `json:"can_change_info"`
    CanInviteUsers bool `json:"can_invite_users"`
    CanPostMessages bool `json:"can_post_messages,omitempty"`
    CanEditMessages bool `json:"can_edit_messages,omitempty"`
    CanPinMessages bool `json:"can_pin_messages,omitempty"`
    CanManageTopics bool `json:"can_manage_topics,omitempty"`
    IsMember bool `json:"is_member"`
    CanSendMessages bool `json:"can_send_messages"`
    CanSendAudios bool `json:"can_send_audios"`
    CanSendDocuments bool `json:"can_send_documents"`
    CanSendPhotos bool `json:"can_send_photos"`
    CanSendVideos bool `json:"can_send_videos"`
    CanSendVideoNotes bool `json:"can_send_video_notes"`
    CanSendVoiceNotes bool `json:"can_send_voice_notes"`
    CanSendPolls bool `json:"can_send_polls"`
    CanSendOtherMessages bool `json:"can_send_other_messages"`
    CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
    UntilDate int64 `json:"until_date"`
}


// Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
    // The member's status in the chat, always "creator"
    Status string `json:"status"`
    // Information about the user
    User *User `json:"user"`
    // True, if the user's presence in the chat is hidden
    IsAnonymous bool `json:"is_anonymous"`
    // Optional. Custom title for this user
    CustomTitle string `json:"custom_title,omitempty"`
}

func (v ChatMemberOwner) GetChatMember() ChatMemberOwner {
    return v
}

// Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
    // The member's status in the chat, always "administrator"
    Status string `json:"status"`
    // Information about the user
    User *User `json:"user"`
    // True, if the bot is allowed to edit administrator privileges of that user
    CanBeEdited bool `json:"can_be_edited"`
    // True, if the user's presence in the chat is hidden
    IsAnonymous bool `json:"is_anonymous"`
    // True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode. Implied by any other administrator privilege
    CanManageChat bool `json:"can_manage_chat"`
    // True, if the administrator can delete messages of other users
    CanDeleteMessages bool `json:"can_delete_messages"`
    // True, if the administrator can manage video chats
    CanManageVideoChats bool `json:"can_manage_video_chats"`
    // True, if the administrator can restrict, ban or unban chat members
    CanRestrictMembers bool `json:"can_restrict_members"`
    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
    CanPromoteMembers bool `json:"can_promote_members"`
    // True, if the user is allowed to change the chat title, photo and other settings
    CanChangeInfo bool `json:"can_change_info"`
    // True, if the user is allowed to invite new users to the chat
    CanInviteUsers bool `json:"can_invite_users"`
    // Optional. True, if the administrator can post in the channel; channels only
    CanPostMessages bool `json:"can_post_messages,omitempty"`
    // Optional. True, if the administrator can edit messages of other users and can pin messages; channels only
    CanEditMessages bool `json:"can_edit_messages,omitempty"`
    // Optional. True, if the user is allowed to pin messages; groups and supergroups only
    CanPinMessages bool `json:"can_pin_messages,omitempty"`
    // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; supergroups only
    CanManageTopics bool `json:"can_manage_topics,omitempty"`
    // Optional. Custom title for this user
    CustomTitle string `json:"custom_title,omitempty"`
}

func (v ChatMemberAdministrator) GetChatMember() ChatMemberAdministrator {
    return v
}

// Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
    // The member's status in the chat, always "member"
    Status string `json:"status"`
    // Information about the user
    User *User `json:"user"`
}

func (v ChatMemberMember) GetChatMember() ChatMemberMember {
    return v
}

// Represents a chat member that is under certain restrictions in the chat. Supergroups only.
type ChatMemberRestricted struct {
    // The member's status in the chat, always "restricted"
    Status string `json:"status"`
    // Information about the user
    User *User `json:"user"`
    // True, if the user is a member of the chat at the moment of the request
    IsMember bool `json:"is_member"`
    // True, if the user is allowed to send text messages, contacts, invoices, locations and venues
    CanSendMessages bool `json:"can_send_messages"`
    // True, if the user is allowed to send audios
    CanSendAudios bool `json:"can_send_audios"`
    // True, if the user is allowed to send documents
    CanSendDocuments bool `json:"can_send_documents"`
    // True, if the user is allowed to send photos
    CanSendPhotos bool `json:"can_send_photos"`
    // True, if the user is allowed to send videos
    CanSendVideos bool `json:"can_send_videos"`
    // True, if the user is allowed to send video notes
    CanSendVideoNotes bool `json:"can_send_video_notes"`
    // True, if the user is allowed to send voice notes
    CanSendVoiceNotes bool `json:"can_send_voice_notes"`
    // True, if the user is allowed to send polls
    CanSendPolls bool `json:"can_send_polls"`
    // True, if the user is allowed to send animations, games, stickers and use inline bots
    CanSendOtherMessages bool `json:"can_send_other_messages"`
    // True, if the user is allowed to add web page previews to their messages
    CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
    // True, if the user is allowed to change the chat title, photo and other settings
    CanChangeInfo bool `json:"can_change_info"`
    // True, if the user is allowed to invite new users to the chat
    CanInviteUsers bool `json:"can_invite_users"`
    // True, if the user is allowed to pin messages
    CanPinMessages bool `json:"can_pin_messages"`
    // True, if the user is allowed to create forum topics
    CanManageTopics bool `json:"can_manage_topics"`
    // Date when restrictions will be lifted for this user; unix time. If 0, then the user is restricted forever
    UntilDate int64 `json:"until_date"`
}

func (v ChatMemberRestricted) GetChatMember() ChatMemberRestricted {
    return v
}

// Represents a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
    // The member's status in the chat, always "left"
    Status string `json:"status"`
    // Information about the user
    User *User `json:"user"`
}

func (v ChatMemberLeft) GetChatMember() ChatMemberLeft {
    return v
}

// Represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
    // The member's status in the chat, always "kicked"
    Status string `json:"status"`
    // Information about the user
    User *User `json:"user"`
    // Date when restrictions will be lifted for this user; unix time. If 0, then the user is banned forever
    UntilDate int64 `json:"until_date"`
}

func (v ChatMemberBanned) GetChatMember() ChatMemberBanned {
    return v
}

// This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
    // Chat the user belongs to
    Chat *Chat `json:"chat"`
    // Performer of the action, which resulted in the change
    From *User `json:"from"`
    // Date the change was done in Unix time
    Date int64 `json:"date"`
    // Previous information about the chat member
    OldChatMember *ChatMember `json:"old_chat_member"`
    // New information about the chat member
    NewChatMember *ChatMember `json:"new_chat_member"`
    // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
    InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
    // Optional. True, if the user joined the chat via a chat folder invite link
    ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}


// Represents a join request sent to a chat.
type ChatJoinRequest struct {
    // Chat to which the request was sent
    Chat *Chat `json:"chat"`
    // User that sent the join request
    From *User `json:"from"`
    // Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot can use this identifier for 24 hours to send messages until the join request is processed, assuming no other administrator contacted the user.
    UserChatId int64 `json:"user_chat_id"`
    // Date the request was sent in Unix time
    Date int64 `json:"date"`
    // Optional. Bio of the user.
    Bio string `json:"bio,omitempty"`
    // Optional. Chat invite link that was used by the user to send the join request
    InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}


// Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
    // Optional. True, if the user is allowed to send text messages, contacts, invoices, locations and venues
    CanSendMessages bool `json:"can_send_messages,omitempty"`
    // Optional. True, if the user is allowed to send audios
    CanSendAudios bool `json:"can_send_audios,omitempty"`
    // Optional. True, if the user is allowed to send documents
    CanSendDocuments bool `json:"can_send_documents,omitempty"`
    // Optional. True, if the user is allowed to send photos
    CanSendPhotos bool `json:"can_send_photos,omitempty"`
    // Optional. True, if the user is allowed to send videos
    CanSendVideos bool `json:"can_send_videos,omitempty"`
    // Optional. True, if the user is allowed to send video notes
    CanSendVideoNotes bool `json:"can_send_video_notes,omitempty"`
    // Optional. True, if the user is allowed to send voice notes
    CanSendVoiceNotes bool `json:"can_send_voice_notes,omitempty"`
    // Optional. True, if the user is allowed to send polls
    CanSendPolls bool `json:"can_send_polls,omitempty"`
    // Optional. True, if the user is allowed to send animations, games, stickers and use inline bots
    CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
    // Optional. True, if the user is allowed to add web page previews to their messages
    CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
    // Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
    CanChangeInfo bool `json:"can_change_info,omitempty"`
    // Optional. True, if the user is allowed to invite new users to the chat
    CanInviteUsers bool `json:"can_invite_users,omitempty"`
    // Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
    CanPinMessages bool `json:"can_pin_messages,omitempty"`
    // Optional. True, if the user is allowed to create forum topics. If omitted defaults to the value of can_pin_messages
    CanManageTopics bool `json:"can_manage_topics,omitempty"`
}


// Represents a location to which a chat is connected.
type ChatLocation struct {
    // The location to which the supergroup is connected. Can't be a live location.
    Location *Location `json:"location"`
    // Location address; 1-64 characters, as defined by the chat owner
    Address string `json:"address"`
}


// This object represents a forum topic.
type ForumTopic struct {
    // Unique identifier of the forum topic
    MessageThreadId int64 `json:"message_thread_id"`
    // Name of the topic
    Name string `json:"name"`
    // Color of the topic icon in RGB format
    IconColor int64 `json:"icon_color"`
    // Optional. Unique identifier of the custom emoji shown as the topic icon
    IconCustomEmojiId string `json:"icon_custom_emoji_id,omitempty"`
}


// This object represents a bot command.
type BotCommand struct {
    // Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
    Command string `json:"command"`
    // Description of the command; 1-256 characters.
    Description string `json:"description"`
}


// This object represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:
// - BotCommandScopeDefault
// - BotCommandScopeAllPrivateChats
// - BotCommandScopeAllGroupChats
// - BotCommandScopeAllChatAdministrators
// - BotCommandScopeChat
// - BotCommandScopeChatAdministrators
// - BotCommandScopeChatMember
type BotCommandScope struct {
    Type string `json:"type"`
    ChatId int64 `json:"chat_id"`
    UserId int64 `json:"user_id"`
}


// Represents the default scope of bot commands. Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
    // Scope type, must be default
    Type string `json:"type"`
}

func (v BotCommandScopeDefault) GetBotCommandScope() BotCommandScopeDefault {
    return v
}

// Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
    // Scope type, must be all_private_chats
    Type string `json:"type"`
}

func (v BotCommandScopeAllPrivateChats) GetBotCommandScope() BotCommandScopeAllPrivateChats {
    return v
}

// Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
    // Scope type, must be all_group_chats
    Type string `json:"type"`
}

func (v BotCommandScopeAllGroupChats) GetBotCommandScope() BotCommandScopeAllGroupChats {
    return v
}

// Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
    // Scope type, must be all_chat_administrators
    Type string `json:"type"`
}

func (v BotCommandScopeAllChatAdministrators) GetBotCommandScope() BotCommandScopeAllChatAdministrators {
    return v
}

// Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
    // Scope type, must be chat
    Type string `json:"type"`
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId int64 `json:"chat_id"`
}

func (v BotCommandScopeChat) GetBotCommandScope() BotCommandScopeChat {
    return v
}

// Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
    // Scope type, must be chat_administrators
    Type string `json:"type"`
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId int64 `json:"chat_id"`
}

func (v BotCommandScopeChatAdministrators) GetBotCommandScope() BotCommandScopeChatAdministrators {
    return v
}

// Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
    // Scope type, must be chat_member
    Type string `json:"type"`
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId int64 `json:"chat_id"`
    // Unique identifier of the target user
    UserId int64 `json:"user_id"`
}

func (v BotCommandScopeChatMember) GetBotCommandScope() BotCommandScopeChatMember {
    return v
}

// This object represents the bot's name.
type BotName struct {
    // The bot's name
    Name string `json:"name"`
}


// This object represents the bot's description.
type BotDescription struct {
    // The bot's description
    Description string `json:"description"`
}


// This object represents the bot's short description.
type BotShortDescription struct {
    // The bot's short description
    ShortDescription string `json:"short_description"`
}


// This object describes the bot's menu button in a private chat. It should be one of
// - MenuButtonCommands
// - MenuButtonWebApp
// - MenuButtonDefault
// If a menu button other than MenuButtonDefault is set for a private chat, then it is applied in the chat. Otherwise the default menu button is applied. By default, the menu button opens the list of bot commands.
type MenuButton struct {
    Type string `json:"type"`
    Text string `json:"text"`
    WebApp *WebAppInfo `json:"web_app"`
}


// Represents a menu button, which opens the bot's list of commands.
type MenuButtonCommands struct {
    // Type of the button, must be commands
    Type string `json:"type"`
}

func (v MenuButtonCommands) GetMenuButton() MenuButtonCommands {
    return v
}

// Represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
    // Type of the button, must be web_app
    Type string `json:"type"`
    // Text on the button
    Text string `json:"text"`
    // Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
    WebApp *WebAppInfo `json:"web_app"`
}

func (v MenuButtonWebApp) GetMenuButton() MenuButtonWebApp {
    return v
}

// Describes that no specific value for the menu button was set.
type MenuButtonDefault struct {
    // Type of the button, must be default
    Type string `json:"type"`
}

func (v MenuButtonDefault) GetMenuButton() MenuButtonDefault {
    return v
}

// Describes why a request was unsuccessful.
type ResponseParameters struct {
    // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
    MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
    // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
    RetryAfter int64 `json:"retry_after,omitempty"`
}


// This object represents the content of a media message to be sent. It should be one of
// - InputMediaAnimation
// - InputMediaDocument
// - InputMediaAudio
// - InputMediaPhoto
// - InputMediaVideo
type InputMedia struct {
    Type string `json:"type"`
    Media string `json:"media"`
    Thumbnail string `json:"thumbnail,omitempty"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    Width int64 `json:"width,omitempty"`
    Height int64 `json:"height,omitempty"`
    Duration int64 `json:"duration,omitempty"`
    HasSpoiler bool `json:"has_spoiler,omitempty"`
    DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
    Performer string `json:"performer,omitempty"`
    Title string `json:"title,omitempty"`
    SupportsStreaming bool `json:"supports_streaming,omitempty"`
}


// Represents a photo to be sent.
type InputMediaPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Media string `json:"media"`
    // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Pass True if the photo needs to be covered with a spoiler animation
    HasSpoiler bool `json:"has_spoiler,omitempty"`
}

func (v InputMediaPhoto) GetInputMedia() InputMediaPhoto {
    return v
}

// Represents a video to be sent.
type InputMediaVideo struct {
    // Type of the result, must be video
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Media string `json:"media"`
    // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Thumbnail string `json:"thumbnail,omitempty"`
    // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Video width
    Width int64 `json:"width,omitempty"`
    // Optional. Video height
    Height int64 `json:"height,omitempty"`
    // Optional. Video duration in seconds
    Duration int64 `json:"duration,omitempty"`
    // Optional. Pass True if the uploaded video is suitable for streaming
    SupportsStreaming bool `json:"supports_streaming,omitempty"`
    // Optional. Pass True if the video needs to be covered with a spoiler animation
    HasSpoiler bool `json:"has_spoiler,omitempty"`
}

func (v InputMediaVideo) GetInputMedia() InputMediaVideo {
    return v
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
    // Type of the result, must be animation
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Media string `json:"media"`
    // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Thumbnail string `json:"thumbnail,omitempty"`
    // Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the animation caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Animation width
    Width int64 `json:"width,omitempty"`
    // Optional. Animation height
    Height int64 `json:"height,omitempty"`
    // Optional. Animation duration in seconds
    Duration int64 `json:"duration,omitempty"`
    // Optional. Pass True if the animation needs to be covered with a spoiler animation
    HasSpoiler bool `json:"has_spoiler,omitempty"`
}

func (v InputMediaAnimation) GetInputMedia() InputMediaAnimation {
    return v
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
    // Type of the result, must be audio
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Media string `json:"media"`
    // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Thumbnail string `json:"thumbnail,omitempty"`
    // Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Duration of the audio in seconds
    Duration int64 `json:"duration,omitempty"`
    // Optional. Performer of the audio
    Performer string `json:"performer,omitempty"`
    // Optional. Title of the audio
    Title string `json:"title,omitempty"`
}

func (v InputMediaAudio) GetInputMedia() InputMediaAudio {
    return v
}

// Represents a general file to be sent.
type InputMediaDocument struct {
    // Type of the result, must be document
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Media string `json:"media"`
    // Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Thumbnail string `json:"thumbnail,omitempty"`
    // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Disables automatic server-side content type detection for files uploaded using multipart/form-data. Always True, if the document is sent as part of an album.
    DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

func (v InputMediaDocument) GetInputMedia() InputMediaDocument {
    return v
}

// This object represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile interface {

}


// This object represents a sticker.
type Sticker struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // Type of the sticker, currently one of "regular", "mask", "custom_emoji". The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
    Type string `json:"type"`
    // Sticker width
    Width int64 `json:"width"`
    // Sticker height
    Height int64 `json:"height"`
    // True, if the sticker is animated
    IsAnimated bool `json:"is_animated"`
    // True, if the sticker is a video sticker
    IsVideo bool `json:"is_video"`
    // Optional. Sticker thumbnail in the .WEBP or .JPG format
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
    // Optional. Emoji associated with the sticker
    Emoji string `json:"emoji,omitempty"`
    // Optional. Name of the sticker set to which the sticker belongs
    SetName string `json:"set_name,omitempty"`
    // Optional. For premium regular stickers, premium animation for the sticker
    PremiumAnimation *File `json:"premium_animation,omitempty"`
    // Optional. For mask stickers, the position where the mask should be placed
    MaskPosition *MaskPosition `json:"mask_position,omitempty"`
    // Optional. For custom emoji stickers, unique identifier of the custom emoji
    CustomEmojiId string `json:"custom_emoji_id,omitempty"`
    // Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
    NeedsRepainting bool `json:"needs_repainting,omitempty"`
    // Optional. File size in bytes
    FileSize int64 `json:"file_size,omitempty"`
}


// This object represents a sticker set.
type StickerSet struct {
    // Sticker set name
    Name string `json:"name"`
    // Sticker set title
    Title string `json:"title"`
    // Type of stickers in the set, currently one of "regular", "mask", "custom_emoji"
    StickerType string `json:"sticker_type"`
    // True, if the sticker set contains animated stickers
    IsAnimated bool `json:"is_animated"`
    // True, if the sticker set contains video stickers
    IsVideo bool `json:"is_video"`
    // List of all set stickers
    Stickers []Sticker `json:"stickers"`
    // Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
    Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}


// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
    // The part of the face relative to which the mask should be placed. One of "forehead", "eyes", "mouth", or "chin".
    Point string `json:"point"`
    // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
    XShift float64 `json:"x_shift"`
    // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
    YShift float64 `json:"y_shift"`
    // Mask scaling coefficient. For example, 2.0 means double size.
    Scale float64 `json:"scale"`
}


// This object describes a sticker to be added to a sticker set.
type InputSticker struct {
    // The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, upload a new one using multipart/form-data, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
    Sticker string `json:"sticker"`
    // List of 1-20 emoji associated with the sticker
    EmojiList []string `json:"emoji_list"`
    // Optional. Position where the mask should be placed on faces. For "mask" stickers only.
    MaskPosition *MaskPosition `json:"mask_position,omitempty"`
    // Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For "regular" and "custom_emoji" stickers only.
    Keywords []string `json:"keywords,omitempty"`
}


// This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
    // Unique identifier for this query
    Id string `json:"id"`
    // Sender
    From *User `json:"from"`
    // Text of the query (up to 256 characters)
    Query string `json:"query"`
    // Offset of the results to be returned, can be controlled by the bot
    Offset string `json:"offset"`
    // Optional. Type of the chat from which the inline query was sent. Can be either "sender" for a private chat with the inline query sender, "private", "group", "supergroup", or "channel". The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
    ChatType string `json:"chat_type,omitempty"`
    // Optional. Sender location, only for bots that request user location
    Location *Location `json:"location,omitempty"`
}


// This object represents a button to be shown above inline query results. You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {
    // Label text on the button
    Text string `json:"text"`
    // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to switch back to the inline mode using the method web_app_switch_inline_query inside the Web App.
    WebApp *WebAppInfo `json:"web_app,omitempty"`
    // Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed. Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
    StartParameter string `json:"start_parameter,omitempty"`
}


// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:
// - InlineQueryResultCachedAudio
// - InlineQueryResultCachedDocument
// - InlineQueryResultCachedGif
// - InlineQueryResultCachedMpeg4Gif
// - InlineQueryResultCachedPhoto
// - InlineQueryResultCachedSticker
// - InlineQueryResultCachedVideo
// - InlineQueryResultCachedVoice
// - InlineQueryResultArticle
// - InlineQueryResultAudio
// - InlineQueryResultContact
// - InlineQueryResultGame
// - InlineQueryResultDocument
// - InlineQueryResultGif
// - InlineQueryResultLocation
// - InlineQueryResultMpeg4Gif
// - InlineQueryResultPhoto
// - InlineQueryResultVenue
// - InlineQueryResultVideo
// - InlineQueryResultVoice
// Note: All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
type InlineQueryResult struct {
    Type string `json:"type"`
    Id string `json:"id"`
    AudioFileId string `json:"audio_file_id"`
    Caption string `json:"caption,omitempty"`
    ParseMode string `json:"parse_mode,omitempty"`
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
    Title string `json:"title"`
    DocumentFileId string `json:"document_file_id"`
    Description string `json:"description,omitempty"`
    GifFileId string `json:"gif_file_id"`
    Mpeg4FileId string `json:"mpeg4_file_id"`
    PhotoFileId string `json:"photo_file_id"`
    StickerFileId string `json:"sticker_file_id"`
    VideoFileId string `json:"video_file_id"`
    VoiceFileId string `json:"voice_file_id"`
    Url string `json:"url,omitempty"`
    HideUrl bool `json:"hide_url,omitempty"`
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`
    ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
    AudioUrl string `json:"audio_url"`
    Performer string `json:"performer,omitempty"`
    AudioDuration int64 `json:"audio_duration,omitempty"`
    PhoneNumber string `json:"phone_number"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`
    Vcard string `json:"vcard,omitempty"`
    GameShortName string `json:"game_short_name"`
    DocumentUrl string `json:"document_url"`
    MimeType string `json:"mime_type"`
    GifUrl string `json:"gif_url"`
    GifWidth int64 `json:"gif_width,omitempty"`
    GifHeight int64 `json:"gif_height,omitempty"`
    GifDuration int64 `json:"gif_duration,omitempty"`
    ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    LivePeriod int64 `json:"live_period,omitempty"`
    Heading int64 `json:"heading,omitempty"`
    ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
    Mpeg4Url string `json:"mpeg4_url"`
    Mpeg4Width int64 `json:"mpeg4_width,omitempty"`
    Mpeg4Height int64 `json:"mpeg4_height,omitempty"`
    Mpeg4Duration int64 `json:"mpeg4_duration,omitempty"`
    PhotoUrl string `json:"photo_url"`
    PhotoWidth int64 `json:"photo_width,omitempty"`
    PhotoHeight int64 `json:"photo_height,omitempty"`
    Address string `json:"address"`
    FoursquareId string `json:"foursquare_id,omitempty"`
    FoursquareType string `json:"foursquare_type,omitempty"`
    GooglePlaceId string `json:"google_place_id,omitempty"`
    GooglePlaceType string `json:"google_place_type,omitempty"`
    VideoUrl string `json:"video_url"`
    VideoWidth int64 `json:"video_width,omitempty"`
    VideoHeight int64 `json:"video_height,omitempty"`
    VideoDuration int64 `json:"video_duration,omitempty"`
    VoiceUrl string `json:"voice_url"`
    VoiceDuration int64 `json:"voice_duration,omitempty"`
}


// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
    // Type of the result, must be article
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Title of the result
    Title string `json:"title"`
    // Content of the message to be sent
    InputMessageContent *InputMessageContent `json:"input_message_content"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. URL of the result
    Url string `json:"url,omitempty"`
    // Optional. Pass True if you don't want the URL to be shown in the message
    HideUrl bool `json:"hide_url,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

func (v InlineQueryResultArticle) GetInlineQueryResult() InlineQueryResultArticle {
    return v
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB
    PhotoUrl string `json:"photo_url"`
    // URL of the thumbnail for the photo
    ThumbnailUrl string `json:"thumbnail_url"`
    // Optional. Width of the photo
    PhotoWidth int64 `json:"photo_width,omitempty"`
    // Optional. Height of the photo
    PhotoHeight int64 `json:"photo_height,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the photo
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultPhoto) GetInlineQueryResult() InlineQueryResultPhoto {
    return v
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
    // Type of the result, must be gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the GIF file. File size must not exceed 1MB
    GifUrl string `json:"gif_url"`
    // Optional. Width of the GIF
    GifWidth int64 `json:"gif_width,omitempty"`
    // Optional. Height of the GIF
    GifHeight int64 `json:"gif_height,omitempty"`
    // Optional. Duration of the GIF in seconds
    GifDuration int64 `json:"gif_duration,omitempty"`
    // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url"`
    // Optional. MIME type of the thumbnail, must be one of "image/jpeg", "image/gif", or "video/mp4". Defaults to "image/jpeg"
    ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the GIF animation
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultGif) GetInlineQueryResult() InlineQueryResultGif {
    return v
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
    // Type of the result, must be mpeg4_gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the MPEG4 file. File size must not exceed 1MB
    Mpeg4Url string `json:"mpeg4_url"`
    // Optional. Video width
    Mpeg4Width int64 `json:"mpeg4_width,omitempty"`
    // Optional. Video height
    Mpeg4Height int64 `json:"mpeg4_height,omitempty"`
    // Optional. Video duration in seconds
    Mpeg4Duration int64 `json:"mpeg4_duration,omitempty"`
    // URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url"`
    // Optional. MIME type of the thumbnail, must be one of "image/jpeg", "image/gif", or "video/mp4". Defaults to "image/jpeg"
    ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video animation
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultMpeg4Gif) GetInlineQueryResult() InlineQueryResultMpeg4Gif {
    return v
}

// Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
    // Type of the result, must be video
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the embedded video player or video file
    VideoUrl string `json:"video_url"`
    // MIME type of the content of the video URL, "text/html" or "video/mp4"
    MimeType string `json:"mime_type"`
    // URL of the thumbnail (JPEG only) for the video
    ThumbnailUrl string `json:"thumbnail_url"`
    // Title for the result
    Title string `json:"title"`
    // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Video width
    VideoWidth int64 `json:"video_width,omitempty"`
    // Optional. Video height
    VideoHeight int64 `json:"video_height,omitempty"`
    // Optional. Video duration in seconds
    VideoDuration int64 `json:"video_duration,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultVideo) GetInlineQueryResult() InlineQueryResultVideo {
    return v
}

// Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultAudio struct {
    // Type of the result, must be audio
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the audio file
    AudioUrl string `json:"audio_url"`
    // Title
    Title string `json:"title"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Performer
    Performer string `json:"performer,omitempty"`
    // Optional. Audio duration in seconds
    AudioDuration int64 `json:"audio_duration,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the audio
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultAudio) GetInlineQueryResult() InlineQueryResultAudio {
    return v
}

// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultVoice struct {
    // Type of the result, must be voice
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the voice recording
    VoiceUrl string `json:"voice_url"`
    // Recording title
    Title string `json:"title"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Recording duration in seconds
    VoiceDuration int64 `json:"voice_duration,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the voice recording
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultVoice) GetInlineQueryResult() InlineQueryResultVoice {
    return v
}

// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultDocument struct {
    // Type of the result, must be document
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // Title for the result
    Title string `json:"title"`
    // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // A valid URL for the file
    DocumentUrl string `json:"document_url"`
    // MIME type of the content of the file, either "application/pdf" or "application/zip"
    MimeType string `json:"mime_type"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the file
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. URL of the thumbnail (JPEG only) for the file
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

func (v InlineQueryResultDocument) GetInlineQueryResult() InlineQueryResultDocument {
    return v
}

// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultLocation struct {
    // Type of the result, must be location
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Location latitude in degrees
    Latitude float64 `json:"latitude"`
    // Location longitude in degrees
    Longitude float64 `json:"longitude"`
    // Location title
    Title string `json:"title"`
    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
    LivePeriod int64 `json:"live_period,omitempty"`
    // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
    Heading int64 `json:"heading,omitempty"`
    // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
    ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the location
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

func (v InlineQueryResultLocation) GetInlineQueryResult() InlineQueryResultLocation {
    return v
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultVenue struct {
    // Type of the result, must be venue
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Latitude of the venue location in degrees
    Latitude float64 `json:"latitude"`
    // Longitude of the venue location in degrees
    Longitude float64 `json:"longitude"`
    // Title of the venue
    Title string `json:"title"`
    // Address of the venue
    Address string `json:"address"`
    // Optional. Foursquare identifier of the venue if known
    FoursquareId string `json:"foursquare_id,omitempty"`
    // Optional. Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
    FoursquareType string `json:"foursquare_type,omitempty"`
    // Optional. Google Places identifier of the venue
    GooglePlaceId string `json:"google_place_id,omitempty"`
    // Optional. Google Places type of the venue. (See supported types.)
    GooglePlaceType string `json:"google_place_type,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the venue
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

func (v InlineQueryResultVenue) GetInlineQueryResult() InlineQueryResultVenue {
    return v
}

// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultContact struct {
    // Type of the result, must be contact
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Contact's phone number
    PhoneNumber string `json:"phone_number"`
    // Contact's first name
    FirstName string `json:"first_name"`
    // Optional. Contact's last name
    LastName string `json:"last_name,omitempty"`
    // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard string `json:"vcard,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the contact
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
    // Optional. Url of the thumbnail for the result
    ThumbnailUrl string `json:"thumbnail_url,omitempty"`
    // Optional. Thumbnail width
    ThumbnailWidth int64 `json:"thumbnail_width,omitempty"`
    // Optional. Thumbnail height
    ThumbnailHeight int64 `json:"thumbnail_height,omitempty"`
}

func (v InlineQueryResultContact) GetInlineQueryResult() InlineQueryResultContact {
    return v
}

// Represents a Game.
// Note: This will only work in Telegram versions released after October 1, 2016. Older clients will not display any inline results if a game result is among them.
type InlineQueryResultGame struct {
    // Type of the result, must be game
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // Short name of the game
    GameShortName string `json:"game_short_name"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (v InlineQueryResultGame) GetInlineQueryResult() InlineQueryResultGame {
    return v
}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier of the photo
    PhotoFileId string `json:"photo_file_id"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the photo caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the photo
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedPhoto) GetInlineQueryResult() InlineQueryResultCachedPhoto {
    return v
}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
    // Type of the result, must be gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the GIF file
    GifFileId string `json:"gif_file_id"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the GIF animation
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedGif) GetInlineQueryResult() InlineQueryResultCachedGif {
    return v
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
    // Type of the result, must be mpeg4_gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the MPEG4 file
    Mpeg4FileId string `json:"mpeg4_file_id"`
    // Optional. Title for the result
    Title string `json:"title,omitempty"`
    // Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video animation
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedMpeg4Gif) GetInlineQueryResult() InlineQueryResultCachedMpeg4Gif {
    return v
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
// Note: This will only work in Telegram versions released after 9 April, 2016 for static stickers and after 06 July, 2019 for animated stickers. Older clients will ignore them.
type InlineQueryResultCachedSticker struct {
    // Type of the result, must be sticker
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier of the sticker
    StickerFileId string `json:"sticker_file_id"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the sticker
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedSticker) GetInlineQueryResult() InlineQueryResultCachedSticker {
    return v
}

// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultCachedDocument struct {
    // Type of the result, must be document
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // Title for the result
    Title string `json:"title"`
    // A valid file identifier for the file
    DocumentFileId string `json:"document_file_id"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the document caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the file
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedDocument) GetInlineQueryResult() InlineQueryResultCachedDocument {
    return v
}

// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
    // Type of the result, must be video
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the video file
    VideoFileId string `json:"video_file_id"`
    // Title for the result
    Title string `json:"title"`
    // Optional. Short description of the result
    Description string `json:"description,omitempty"`
    // Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the video caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the video
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedVideo) GetInlineQueryResult() InlineQueryResultCachedVideo {
    return v
}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultCachedVoice struct {
    // Type of the result, must be voice
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the voice message
    VoiceFileId string `json:"voice_file_id"`
    // Voice message title
    Title string `json:"title"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the voice message caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the voice message
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedVoice) GetInlineQueryResult() InlineQueryResultCachedVoice {
    return v
}

// Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultCachedAudio struct {
    // Type of the result, must be audio
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the audio file
    AudioFileId string `json:"audio_file_id"`
    // Optional. Caption, 0-1024 characters after entities parsing
    Caption string `json:"caption,omitempty"`
    // Optional. Mode for parsing entities in the audio caption. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in the caption, which can be specified instead of parse_mode
    CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
    // Optional. Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Optional. Content of the message to be sent instead of the audio
    InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
}

func (v InlineQueryResultCachedAudio) GetInlineQueryResult() InlineQueryResultCachedAudio {
    return v
}

// This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:
// - InputTextMessageContent
// - InputLocationMessageContent
// - InputVenueMessageContent
// - InputContactMessageContent
// - InputInvoiceMessageContent
type InputMessageContent struct {
    MessageText string `json:"message_text"`
    ParseMode string `json:"parse_mode,omitempty"`
    Entities []MessageEntity `json:"entities,omitempty"`
    DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    LivePeriod int64 `json:"live_period,omitempty"`
    Heading int64 `json:"heading,omitempty"`
    ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
    Title string `json:"title"`
    Address string `json:"address"`
    FoursquareId string `json:"foursquare_id,omitempty"`
    FoursquareType string `json:"foursquare_type,omitempty"`
    GooglePlaceId string `json:"google_place_id,omitempty"`
    GooglePlaceType string `json:"google_place_type,omitempty"`
    PhoneNumber string `json:"phone_number"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name,omitempty"`
    Vcard string `json:"vcard,omitempty"`
    Description string `json:"description"`
    Payload string `json:"payload"`
    ProviderToken string `json:"provider_token"`
    Currency string `json:"currency"`
    Prices []LabeledPrice `json:"prices"`
    MaxTipAmount int64 `json:"max_tip_amount,omitempty"`
    SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`
    ProviderData string `json:"provider_data,omitempty"`
    PhotoUrl string `json:"photo_url,omitempty"`
    PhotoSize int64 `json:"photo_size,omitempty"`
    PhotoWidth int64 `json:"photo_width,omitempty"`
    PhotoHeight int64 `json:"photo_height,omitempty"`
    NeedName bool `json:"need_name,omitempty"`
    NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
    NeedEmail bool `json:"need_email,omitempty"`
    NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
    SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
    SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
    IsFlexible bool `json:"is_flexible,omitempty"`
}


// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
    // Text of the message to be sent, 1-4096 characters
    MessageText string `json:"message_text"`
    // Optional. Mode for parsing entities in the message text. See formatting options for more details.
    ParseMode string `json:"parse_mode,omitempty"`
    // Optional. List of special entities that appear in message text, which can be specified instead of parse_mode
    Entities []MessageEntity `json:"entities,omitempty"`
    // Optional. Disables link previews for links in the sent message
    DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

func (v InputTextMessageContent) GetInputMessageContent() InputTextMessageContent {
    return v
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
    // Latitude of the location in degrees
    Latitude float64 `json:"latitude"`
    // Longitude of the location in degrees
    Longitude float64 `json:"longitude"`
    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
    HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
    // Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
    LivePeriod int64 `json:"live_period,omitempty"`
    // Optional. For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
    Heading int64 `json:"heading,omitempty"`
    // Optional. For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
    ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
}

func (v InputLocationMessageContent) GetInputMessageContent() InputLocationMessageContent {
    return v
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
    // Latitude of the venue in degrees
    Latitude float64 `json:"latitude"`
    // Longitude of the venue in degrees
    Longitude float64 `json:"longitude"`
    // Name of the venue
    Title string `json:"title"`
    // Address of the venue
    Address string `json:"address"`
    // Optional. Foursquare identifier of the venue, if known
    FoursquareId string `json:"foursquare_id,omitempty"`
    // Optional. Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
    FoursquareType string `json:"foursquare_type,omitempty"`
    // Optional. Google Places identifier of the venue
    GooglePlaceId string `json:"google_place_id,omitempty"`
    // Optional. Google Places type of the venue. (See supported types.)
    GooglePlaceType string `json:"google_place_type,omitempty"`
}

func (v InputVenueMessageContent) GetInputMessageContent() InputVenueMessageContent {
    return v
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
    // Contact's phone number
    PhoneNumber string `json:"phone_number"`
    // Contact's first name
    FirstName string `json:"first_name"`
    // Optional. Contact's last name
    LastName string `json:"last_name,omitempty"`
    // Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard string `json:"vcard,omitempty"`
}

func (v InputContactMessageContent) GetInputMessageContent() InputContactMessageContent {
    return v
}

// Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
    // Product name, 1-32 characters
    Title string `json:"title"`
    // Product description, 1-255 characters
    Description string `json:"description"`
    // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
    Payload string `json:"payload"`
    // Payment provider token, obtained via @BotFather
    ProviderToken string `json:"provider_token"`
    // Three-letter ISO 4217 currency code, see more on currencies
    Currency string `json:"currency"`
    // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
    Prices []LabeledPrice `json:"prices"`
    // Optional. The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0
    MaxTipAmount int64 `json:"max_tip_amount,omitempty"`
    // Optional. A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
    SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"`
    // Optional. A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
    ProviderData string `json:"provider_data,omitempty"`
    // Optional. URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
    PhotoUrl string `json:"photo_url,omitempty"`
    // Optional. Photo size in bytes
    PhotoSize int64 `json:"photo_size,omitempty"`
    // Optional. Photo width
    PhotoWidth int64 `json:"photo_width,omitempty"`
    // Optional. Photo height
    PhotoHeight int64 `json:"photo_height,omitempty"`
    // Optional. Pass True if you require the user's full name to complete the order
    NeedName bool `json:"need_name,omitempty"`
    // Optional. Pass True if you require the user's phone number to complete the order
    NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
    // Optional. Pass True if you require the user's email address to complete the order
    NeedEmail bool `json:"need_email,omitempty"`
    // Optional. Pass True if you require the user's shipping address to complete the order
    NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
    // Optional. Pass True if the user's phone number should be sent to provider
    SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
    // Optional. Pass True if the user's email address should be sent to provider
    SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
    // Optional. Pass True if the final price depends on the shipping method
    IsFlexible bool `json:"is_flexible,omitempty"`
}

func (v InputInvoiceMessageContent) GetInputMessageContent() InputInvoiceMessageContent {
    return v
}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// Note: It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
type ChosenInlineResult struct {
    // The unique identifier for the result that was chosen
    ResultId string `json:"result_id"`
    // The user that chose the result
    From *User `json:"from"`
    // Optional. Sender location, only for bots that require user location
    Location *Location `json:"location,omitempty"`
    // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
    InlineMessageId string `json:"inline_message_id,omitempty"`
    // The query that was used to obtain the result
    Query string `json:"query"`
}


// Describes an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
    // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message.
    InlineMessageId string `json:"inline_message_id,omitempty"`
}


// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
    // Portion label
    Label string `json:"label"`
    // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    Amount int64 `json:"amount"`
}


// This object contains basic information about an invoice.
type Invoice struct {
    // Product name
    Title string `json:"title"`
    // Product description
    Description string `json:"description"`
    // Unique bot deep-linking parameter that can be used to generate this invoice
    StartParameter string `json:"start_parameter"`
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int64 `json:"total_amount"`
}


// This object represents a shipping address.
type ShippingAddress struct {
    // Two-letter ISO 3166-1 alpha-2 country code
    CountryCode string `json:"country_code"`
    // State, if applicable
    State string `json:"state"`
    // City
    City string `json:"city"`
    // First line for the address
    StreetLine1 string `json:"street_line1"`
    // Second line for the address
    StreetLine2 string `json:"street_line2"`
    // Address post code
    PostCode string `json:"post_code"`
}


// This object represents information about an order.
type OrderInfo struct {
    // Optional. User name
    Name string `json:"name,omitempty"`
    // Optional. User's phone number
    PhoneNumber string `json:"phone_number,omitempty"`
    // Optional. User email
    Email string `json:"email,omitempty"`
    // Optional. User shipping address
    ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}


// This object represents one shipping option.
type ShippingOption struct {
    // Shipping option identifier
    Id string `json:"id"`
    // Option title
    Title string `json:"title"`
    // List of price portions
    Prices []LabeledPrice `json:"prices"`
}


// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int64 `json:"total_amount"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload"`
    // Optional. Identifier of the shipping option chosen by the user
    ShippingOptionId string `json:"shipping_option_id,omitempty"`
    // Optional. Order information provided by the user
    OrderInfo *OrderInfo `json:"order_info,omitempty"`
    // Telegram payment identifier
    TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
    // Provider payment identifier
    ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}


// This object contains information about an incoming shipping query.
type ShippingQuery struct {
    // Unique query identifier
    Id string `json:"id"`
    // User who sent the query
    From *User `json:"from"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload"`
    // User specified shipping address
    ShippingAddress *ShippingAddress `json:"shipping_address"`
}


// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
    // Unique query identifier
    Id string `json:"id"`
    // User who sent the query
    From *User `json:"from"`
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int64 `json:"total_amount"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload"`
    // Optional. Identifier of the shipping option chosen by the user
    ShippingOptionId string `json:"shipping_option_id,omitempty"`
    // Optional. Order information provided by the user
    OrderInfo *OrderInfo `json:"order_info,omitempty"`
}


// Describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
    // Array with information about documents and other Telegram Passport elements that was shared with the bot
    Data []EncryptedPassportElement `json:"data"`
    // Encrypted credentials required to decrypt the data
    Credentials *EncryptedCredentials `json:"credentials"`
}


// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
    // Identifier for this file, which can be used to download or reuse the file
    FileId string `json:"file_id"`
    // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
    FileUniqueId string `json:"file_unique_id"`
    // File size in bytes
    FileSize int64 `json:"file_size"`
    // Unix time when the file was uploaded
    FileDate int64 `json:"file_date"`
}


// Describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
    // Element type. One of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration", "phone_number", "email".
    Type string `json:"type"`
    // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for "personal_details", "passport", "driver_license", "identity_card", "internal_passport" and "address" types. Can be decrypted and verified using the accompanying EncryptedCredentials.
    Data string `json:"data,omitempty"`
    // Optional. User's verified phone number, available only for "phone_number" type
    PhoneNumber string `json:"phone_number,omitempty"`
    // Optional. User's verified email address, available only for "email" type
    Email string `json:"email,omitempty"`
    // Optional. Array of encrypted files with documents provided by the user, available for "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
    Files []PassportFile `json:"files,omitempty"`
    // Optional. Encrypted file with the front side of the document, provided by the user. Available for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
    FrontSide *PassportFile `json:"front_side,omitempty"`
    // Optional. Encrypted file with the reverse side of the document, provided by the user. Available for "driver_license" and "identity_card". The file can be decrypted and verified using the accompanying EncryptedCredentials.
    ReverseSide *PassportFile `json:"reverse_side,omitempty"`
    // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for "passport", "driver_license", "identity_card" and "internal_passport". The file can be decrypted and verified using the accompanying EncryptedCredentials.
    Selfie *PassportFile `json:"selfie,omitempty"`
    // Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration" and "temporary_registration" types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
    Translation []PassportFile `json:"translation,omitempty"`
    // Base64-encoded element hash for using in PassportElementErrorUnspecified
    Hash string `json:"hash"`
}


// Describes data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
    // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
    Data string `json:"data"`
    // Base64-encoded data hash for data authentication
    Hash string `json:"hash"`
    // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
    Secret string `json:"secret"`
}


// This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:
// - PassportElementErrorDataField
// - PassportElementErrorFrontSide
// - PassportElementErrorReverseSide
// - PassportElementErrorSelfie
// - PassportElementErrorFile
// - PassportElementErrorFiles
// - PassportElementErrorTranslationFile
// - PassportElementErrorTranslationFiles
// - PassportElementErrorUnspecified
type PassportElementError struct {
    Source string `json:"source"`
    Type string `json:"type"`
    FieldName string `json:"field_name"`
    DataHash string `json:"data_hash"`
    Message string `json:"message"`
    FileHash string `json:"file_hash"`
    FileHashes []string `json:"file_hashes"`
    ElementHash string `json:"element_hash"`
}


// Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
    // Error source, must be data
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the error, one of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address"
    Type string `json:"type"`
    // Name of the data field which has the error
    FieldName string `json:"field_name"`
    // Base64-encoded data hash
    DataHash string `json:"data_hash"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorDataField) GetPassportElementError() PassportElementErrorDataField {
    return v
}

// Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
    // Error source, must be front_side
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport"
    Type string `json:"type"`
    // Base64-encoded hash of the file with the front side of the document
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorFrontSide) GetPassportElementError() PassportElementErrorFrontSide {
    return v
}

// Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
    // Error source, must be reverse_side
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of "driver_license", "identity_card"
    Type string `json:"type"`
    // Base64-encoded hash of the file with the reverse side of the document
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorReverseSide) GetPassportElementError() PassportElementErrorReverseSide {
    return v
}

// Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
    // Error source, must be selfie
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport"
    Type string `json:"type"`
    // Base64-encoded hash of the file with the selfie
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorSelfie) GetPassportElementError() PassportElementErrorSelfie {
    return v
}

// Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
    // Error source, must be file
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
    Type string `json:"type"`
    // Base64-encoded file hash
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorFile) GetPassportElementError() PassportElementErrorFile {
    return v
}

// Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
    // Error source, must be files
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
    Type string `json:"type"`
    // List of base64-encoded file hashes
    FileHashes []string `json:"file_hashes"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorFiles) GetPassportElementError() PassportElementErrorFiles {
    return v
}

// Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
    // Error source, must be translation_file
    Source string `json:"source"`
    // Type of element of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
    Type string `json:"type"`
    // Base64-encoded file hash
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorTranslationFile) GetPassportElementError() PassportElementErrorTranslationFile {
    return v
}

// Represents an issue with the translated version of a document. The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
    // Error source, must be translation_files
    Source string `json:"source"`
    // Type of element of the user's Telegram Passport which has the issue, one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
    Type string `json:"type"`
    // List of base64-encoded file hashes
    FileHashes []string `json:"file_hashes"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorTranslationFiles) GetPassportElementError() PassportElementErrorTranslationFiles {
    return v
}

// Represents an issue in an unspecified place. The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
    // Error source, must be unspecified
    Source string `json:"source"`
    // Type of element of the user's Telegram Passport which has the issue
    Type string `json:"type"`
    // Base64-encoded element hash
    ElementHash string `json:"element_hash"`
    // Error message
    Message string `json:"message"`
}

func (v PassportElementErrorUnspecified) GetPassportElementError() PassportElementErrorUnspecified {
    return v
}

// This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
    // Title of the game
    Title string `json:"title"`
    // Description of the game
    Description string `json:"description"`
    // Photo that will be displayed in the game message in chats.
    Photo []PhotoSize `json:"photo"`
    // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
    Text string `json:"text,omitempty"`
    // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
    TextEntities []MessageEntity `json:"text_entities,omitempty"`
    // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
    Animation *Animation `json:"animation,omitempty"`
}


// A placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame interface {

}


// This object represents one row of the high scores table for a game.
type GameHighScore struct {
    // Position in high score table for the game
    Position int64 `json:"position"`
    // User
    User *User `json:"user"`
    // Score
    Score int64 `json:"score"`
}


type ReplyMarkup interface {
    replyMarkup()
}
// Unmarshal MessageEntity json arrays
func UnmarshalMessageEntityArray(r json.RawMessage) (*[]MessageEntity, error) {
    var tmp *[]MessageEntity
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal PhotoSize json arrays
func UnmarshalPhotoSizeArray(r json.RawMessage) (*[]PhotoSize, error) {
    var tmp *[]PhotoSize
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal User json arrays
func UnmarshalUserArray(r json.RawMessage) (*[]User, error) {
    var tmp *[]User
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal PollOption json arrays
func UnmarshalPollOptionArray(r json.RawMessage) (*[]PollOption, error) {
    var tmp *[]PollOption
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal Sticker json arrays
func UnmarshalStickerArray(r json.RawMessage) (*[]Sticker, error) {
    var tmp *[]Sticker
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal LabeledPrice json arrays
func UnmarshalLabeledPriceArray(r json.RawMessage) (*[]LabeledPrice, error) {
    var tmp *[]LabeledPrice
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal EncryptedPassportElement json arrays
func UnmarshalEncryptedPassportElementArray(r json.RawMessage) (*[]EncryptedPassportElement, error) {
    var tmp *[]EncryptedPassportElement
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal PassportFile json arrays
func UnmarshalPassportFileArray(r json.RawMessage) (*[]PassportFile, error) {
    var tmp *[]PassportFile
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal PhotoSize json bi-dimensional arrays
func UnmarshalPhotoSizeArrayOfArray(r json.RawMessage) (*[][]PhotoSize, error) {
    var tmp *[][]PhotoSize
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal KeyboardButton json bi-dimensional arrays
func UnmarshalKeyboardButtonArrayOfArray(r json.RawMessage) (*[][]KeyboardButton, error) {
    var tmp *[][]KeyboardButton
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

// Unmarshal InlineKeyboardButton json bi-dimensional arrays
func UnmarshalInlineKeyboardButtonArrayOfArray(r json.RawMessage) (*[][]InlineKeyboardButton, error) {
    var tmp *[][]InlineKeyboardButton
    err := json.Unmarshal(r, &tmp)
    if err != nil {
        return nil, err
    }
    return tmp, err
}

