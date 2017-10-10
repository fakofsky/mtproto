package mtproto

func (m *MTProto) MessagesGetHistory(peer TL, offsetId, offsetDate, addOffset, limit, maxId, minId int32) (*TL, error) {
	return m.InvokeSync(TL_messages_getHistory{
		Peer:        peer,
		Offset_id:   offsetId,
		Offset_date: offsetDate,
		Add_offset:  addOffset,
		Limit:       limit,
		Max_id:      maxId,
		Min_id:      minId,
	})
}

func (m *MTProto) MessagesGetMessagesViews(peer TL,id []int32, incr TL) (*TL, error) {
	return m.InvokeSync(TL_messages_getMessagesViews{
		Peer: peer,
		Id: id,
		Increment: incr,
	})
}

func (m *MTProto) ChannelsGetMessages(channel TL,id []int32) (*TL, error) {
	return m.InvokeSync(TL_channels_getMessages{
		Channel: channel,
		Id: id,
	})
}

func (m *MTProto) MessagesGetAllChats(except_ids []int32) (*TL, error) {
	return m.InvokeSync(TL_messages_getAllChats{
		Except_ids: except_ids,
	})
}

func (m *MTProto) MessagesGetDialogs(excludePinned bool, offsetDate, offsetId int32, offsetPeer TL, limit int32) (*TL, error) {
	return m.InvokeSync(TL_messages_getDialogs{
		Exclude_pinned: excludePinned,
		Offset_date:    offsetDate,
		Offset_id:      offsetId,
		Offset_peer:    offsetPeer,
		Limit:          limit,
	})
}

func (m *MTProto) MessagesSendMessage(no_webpage, silent, background, clear_draft bool, peer TL, reply_to_msg_id int32, message string, random_id int64, reply_markup TL, entities []TL) (*TL, error) {
	return m.InvokeSync(TL_messages_sendMessage{
		No_webpage:      no_webpage,
		Silent:          silent,
		Background:      background,
		Clear_draft:     clear_draft,
		Peer:            peer,
		Reply_to_msg_id: reply_to_msg_id,
		Message:         message,
		Random_id:       random_id,
		Reply_markup:    reply_markup,
		Entities:        entities,
	})
}

func (m *MTProto) MessagesForwardMessages(flags int32,silent bool, background bool,with_my_score bool, from_peer TL, id []int32, random_id []int64, to_peer TL) (*TL,error) {
	return m.InvokeSync(TL_messages_forwardMessages{
		Flags: flags,
		Silent: silent,
		Background: background,
		With_my_score: with_my_score,
		From_peer: from_peer,
		Id: id,
		Random_id: random_id,
		To_peer: to_peer,
	})
}
