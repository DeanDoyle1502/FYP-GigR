import { useEffect, useState } from "react";
import axios from "axios";

const CURRENT_USER_ID = 5;

type Message = {
  senderId: number;
  receiverId: number;
  content: string;
  timestamp: string;
};

type ChatSession = {
  id: string;
  gigId: number;
  userA: number;
  userB: number;
  completedBy: { [userId: number]: boolean };
  isArchived: boolean;
  createdAt: string;
};

type ChatViewProps = {
  gigID: number;
  otherUserID: number;
};

const ChatView: React.FC<ChatViewProps> = ({ gigID, otherUserID }) => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [session, setSession] = useState<ChatSession | null>(null);
  const [newMessage, setNewMessage] = useState("");
  const [loading, setLoading] = useState(true);

  const fetchChat = async () => {
    const token = localStorage.getItem("token");

    try {
      const [sessionRes, messagesRes] = await Promise.all([
        axios.get(`http://localhost:8080/gigs/${gigID}/session/${otherUserID}`, {
          headers: { Authorization: `Bearer ${token}` },
        }),
        axios.get(`http://localhost:8080/gigs/${gigID}/thread/${otherUserID}`, {
          headers: { Authorization: `Bearer ${token}` },
        }),
      ]);
      setSession(sessionRes.data);
      setMessages(messagesRes.data);
    } catch (err) {
      console.error("Failed to load chat", err);
    } finally {
      setLoading(false);
    }
  };

  const sendMessage = async () => {
    const token = localStorage.getItem("token");

    try {
      await axios.post(
        `http://localhost:8080/gigs/${gigID}/messages`,
        {
          receiverId: otherUserID,
          content: newMessage,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      setNewMessage("");
      fetchChat(); // refresh messages
    } catch (err) {
      console.error("Failed to send message", err);
    }
  };

  const markComplete = async () => {
    const token = localStorage.getItem("token");

    try {
      await axios.patch(
        `http://localhost:8080/gigs/${gigID}/session/${otherUserID}/complete`,
        {},
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      fetchChat(); // refresh session state
    } catch (err) {
      console.error("Failed to complete chat", err);
    }
  };

  useEffect(() => {
    fetchChat();
  }, []);

  if (loading) return <div>Loading chat...</div>;

  return (
    <div>
      <h2>Chat for Gig #{gigID}</h2>

      {session?.isArchived && <p>📁 This chat is archived</p>}

      <div className="message-thread" style={{ marginBottom: "1rem" }}>
        {messages.map((msg) => (
          <div key={msg.timestamp}>
            <strong>{msg.senderId === CURRENT_USER_ID ? "You" : `User ${msg.senderId}`}</strong>:{" "}
            {msg.content}
          </div>
        ))}
      </div>

      {!session?.isArchived && (
        <>
          <textarea
            value={newMessage}
            onChange={(e) => setNewMessage(e.target.value)}
            placeholder="Type your message..."
            style={{ width: "100%", height: "80px", marginBottom: "0.5rem" }}
          />
          <button onClick={sendMessage} disabled={!newMessage.trim()}>
            Send
          </button>
        </>
      )}

      <br />

      <button
        onClick={markComplete}
        disabled={session?.completedBy?.[CURRENT_USER_ID]}
        style={{ marginTop: "1rem" }}
      >
        ✅ Mark as Completed
      </button>
    </div>
  );
};

export default ChatView;
