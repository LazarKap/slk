package help

// Sheet is the task-oriented cheat sheet the ? overlay shows: what you
// want to do, grouped, with the key or command that does it. Keep the
// labels in step with keys.go and the slash/colon command tables.
func Sheet() []Entry {
	var out []Entry
	section := func(name string, rows ...[2]string) {
		out = append(out, Entry{Section: name, Header: true})
		for _, r := range rows {
			out = append(out, Entry{Section: name, Key: r[0], Desc: r[1]})
		}
	}

	section("Basics",
		[2]string{"j/k", "move down / up"},
		[2]string{"h/l, tab", "switch panel (sidebar, chat, thread)"},
		[2]string{"enter", "open what is selected"},
		[2]string{"esc", "back / close"},
		[2]string{"i", "insert mode: write a message"},
		[2]string{":", "command mode (:ws, :sp, :marks ...)"},
		[2]string{"?", "this sheet; / inside it filters"},
		[2]string{"Q / ctrl+c", "quit"},
	)
	section("Channels",
		[2]string{"ctrl+t", "find any channel or DM; ctrl+n / ctrl+p to move"},
		[2]string{"/open #name", "open or join a channel (/join alias)"},
		[2]string{"a / A", "next / previous channel with unread"},
		[2]string{"ctrl+h / ctrl+l", "back / forward through visited channels"},
		[2]string{"space", "collapse / expand a sidebar section"},
		[2]string{"] / [", "widen / narrow the sidebar"},
		[2]string{"ctrl+b", "hide / show the sidebar"},
		[2]string{"U", "mark channel unread from the selected message"},
		[2]string{"/leave", "not supported in slk yet; use Slack"},
	)
	section("People and DMs",
		[2]string{"/msg @name text", "DM anyone; @ opens the people picker (/dm alias)"},
		[2]string{"ctrl+t", "jump to a DM you already have"},
		[2]string{"K", "show the author's profile"},
		[2]string{"ctrl+n", "new message picker"},
	)
	section("Read messages",
		[2]string{"gg / G", "oldest loaded / newest"},
		[2]string{"ctrl+u / ctrl+d", "half page up / down (tmux may eat ctrl+d)"},
		[2]string{"y", "copy message text"},
		[2]string{"Y / C", "copy permalink"},
		[2]string{"o", "open a link in the message"},
		[2]string{"O / v", "preview image; enter opens it in Preview.app"},
		[2]string{"d", "download file or image to ~/Downloads"},
		[2]string{"L", "who reacted"},
		[2]string{"b", "press a button in a bot message"},
		[2]string{"x", "remove a link preview"},
	)
	section("Write, edit, delete",
		[2]string{"i ... enter", "write, then send"},
		[2]string{"shift+enter", "new line in a message"},
		[2]string{"@ / :", "mention picker / emoji picker while typing"},
		[2]string{"ctrl+e", "edit the draft in nvim (:wq sends it back)"},
		[2]string{"ctrl+v", "paste image, file path or text as attachment"},
		[2]string{"ctrl+u", "clear the draft (insert mode)"},
		[2]string{"E", "edit your message"},
		[2]string{"D", "delete your message (confirms)"},
		[2]string{"F", "forward a message to another channel"},
	)
	section("Threads",
		[2]string{"enter", "open the thread of the selected message"},
		[2]string{"ctrl+]", "show / hide the thread panel"},
		[2]string{"q", "close the thread"},
		[2]string{"ctrl+o", "reply also to the channel (toggle, in thread compose)"},
		[2]string{"alt+enter", "send this reply to the channel too"},
		[2]string{"S", "save the thread as markdown"},
	)
	section("Reactions",
		[2]string{"r", "add a reaction (type to search, arrows to pick)"},
		[2]string{"R", "toggle / navigate existing reactions"},
		[2]string{"L", "list who reacted"},
	)
	section("Search",
		[2]string{"/", "search in this channel (loaded history), n / N next / previous"},
		[2]string{"ctrl+f", "search the workspace"},
		[2]string{"from:@user", "workspace search: by author"},
		[2]string{"in:#channel", "workspace search: in a channel"},
		[2]string{"before:2026-01-31", "workspace search: by date (after: too)"},
		[2]string{"esc", "clear the channel search"},
	)
	section("Activity (ctrl+a)",
		[2]string{"ctrl+a / :activity", "open / close the activity view"},
		[2]string{"f", "filter: all, mentions, threads, DMs, reactions"},
		[2]string{"p", "priority sort: unread mentions and DMs first"},
		[2]string{"u", "unread only"},
		[2]string{"U", "mark one item unread / read"},
		[2]string{"R", "mark everything read"},
		[2]string{"enter", "jump to the message or thread"},
	)
	section("Marks (jump points)",
		[2]string{"m a-z", "set a mark for this session"},
		[2]string{"m A-Z", "set a mark that survives restarts"},
		[2]string{"' x", "jump to mark x"},
		[2]string{"''", "jump back to where you came from"},
		[2]string{":marks", "list marks"},
		[2]string{":delmarks abc", "delete marks"},
	)
	section("Windows (splits)",
		[2]string{"ctrl+w s / :sp", "split horizontally"},
		[2]string{"ctrl+w v / :vsp", "split vertically"},
		[2]string{"ctrl+w h/j/k/l", "focus the window in that direction"},
		[2]string{"ctrl+w w", "cycle windows"},
		[2]string{"ctrl+w q / :q", "close this window"},
		[2]string{"ctrl+w o / :only", "close the other windows"},
	)
	section("Workspace, status, look",
		[2]string{"1-9", "jump to workspace N"},
		[2]string{":ws", "workspace picker"},
		[2]string{"ctrl+g / ctrl+s", "presence and status: active, away, snooze, WFH"},
		[2]string{"/active, /away", "set presence from the composer"},
		[2]string{"/status", "not supported in slk yet"},
		[2]string{"ctrl+y", "pick a theme (saved per workspace)"},
		[2]string{"alt+y / alt+shift+y", "next / previous theme"},
		[2]string{"ctrl+shift+y", "set the default theme"},
		[2]string{"H / :huddle", "copy the huddle link"},
		[2]string{":help", "this sheet"},
		[2]string{":version", "show the running build"},
	)
	return out
}
