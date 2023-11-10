// composables/useMessage.js
import { ref, onMounted } from 'vue';

export default function useMessage() {
    const messages = ref([]);
    let messageContainer = null;
    const messageType = {
        info: 'rounded-md px-4 py-2 mb-2',
        success: 'rounded-md px-4 py-2 mb-2',
        warning: 'rounded-md px-4 py-2 mb-2',
        error: 'rounded-md px-4 py-2 mb-2',

    }
    // 初始化消息容器
    onMounted(() => {
        messageContainer = document.createElement('div');
        messageContainer.className = 'fixed top-20 left-1/2 transform -translate-x-1/2 z-[9999]'
        document.body.appendChild(messageContainer);
    });

    // 添加新消息
    const addMessage = (text, type = 'info', duration = 3000) => {
        const id = Date.now();
        const messageElement = document.createElement('div');
        messageElement.className = 'shadow-center ' + messageType[type];
        messageElement.textContent = text;
        messageElement.id = `message-${id}`;
        messageContainer.appendChild(messageElement);

        // 将消息添加到响应式数组
        messages.value.push({ id, text, type });

        // 自动删除消息
        setTimeout(() => {
            removeMessage(id);
        }, duration);
    };

    // 删除指定消息
    const removeMessage = (id) => {
        const index = messages.value.findIndex(message => message.id === id);
        if (index > -1) {
            // 从 DOM 中移除
            const messageElement = document.getElementById(`message-${id}`);
            messageContainer.removeChild(messageElement);

            // 从响应式数组中移除
            messages.value.splice(index, 1);
        }
    };

    return {
        messages,
        addMessage,
        removeMessage
    };
}
