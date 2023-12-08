// composables/useWebSocket.js
import { useMessageStore } from "~/stores/message.js";
export default function useWebSocket() {
    const messageStore = useMessageStore();
    const { getCookie } = useCookies();
    const { notice } = useNotice();
    let socket = null;
    let reconnectAttempts = 0;
    const maxReconnectAttempts = 5; // 最大重连次数
    const reconnectInterval = 5000; // 重连间隔时间（毫秒）

    onBeforeUnmount(() => {
        closeInterval();
    });



    let interval = null;
    function initSocket() {
        let token = getCookie("token");
        socket = new WebSocket(`ws://localhost:8080/ws?Authorization=${token}`);
        socket.onopen = handleSocketOpen;
        socket.onmessage = handleSocketMessage;
        socket.onclose = handleSocketClose;
        socket.onerror = handleSocketError;
    }
    const handleSocketMessage = (data) => {
        try {
            let message = JSON.parse(data.data);
            if (!message.type) return;
            messageStore.setUnreadMessage(messageStore.getUnreadMessage + 1);
            notice({ content: `${message.data}`, })
        } catch (error) { }
    };
    function handleSocketOpen(event) {
        heartCheck();
        interval = setInterval(heartCheck, 10000);
        reconnectAttempts = 0; // 重置重连尝试次数
    }

    function handleSocketClose(event) {
        console.log("Client notified socket has closed", event);
        tryReconnect(); // 尝试重连
    }

    function handleSocketError(event) {
        console.log("Client notified socket has errored", event);
        tryReconnect(); // 尝试重连
    }

    function heartCheck() {
        socket.send(JSON.stringify({ ping: new Date().getTime() }));
    }

    function tryReconnect() {
        if (reconnectAttempts < maxReconnectAttempts) {
            setTimeout(() => {
                reconnectAttempts++;
                initSocket();
            }, reconnectInterval);
        }
    }
    function closeInterval() {
        if (interval) {
            clearInterval(interval);
        }
    }
    function closeSocket() {
        if (socket) {
            socket.close();
        }
    }
    return { initSocket, closeSocket };
}
