<template>
  <NuxtLayout>
    <NuxtLoadingIndicator class="!transform-none top-16" />
    <NuxtPage />
  </NuxtLayout>
  <Teleport to="body">
    <FloatingList />
  </Teleport>
</template>
<script setup>
const { getCookie } = useCookies();
import { useAppConfigStore } from "~/stores/init";
import { useUserStore } from "~/stores/main";
let app = useAppConfigStore();
app.fetchHost();
let userStore = useUserStore();
let isLogin = computed(() => userStore.isLogin);
let socket = null;

watch(
  () => isLogin.value,
  (val) => {
    if (val) {
      console.log(val);
      initSocket();
    } else {
      if (socket) {
        socket.close();
      }
    }
  },
  { immediate: true }
);
function initSocket() {
  let token = getCookie('token')
  // 创建一个 WebSocket 连接
  socket = new WebSocket(`ws://localhost:8080/ws?Authorization=${token}`);
  // 打开 WebSocket 连接成功后，发送消息
  socket.onopen = function (event) {
    socket.send("Hello Server!");
  };
  // 接收到服务端发送的消息后，打印出来
  socket.onmessage = function (event) {
    console.log("Client received a message", event);
  };
  // 关闭 WebSocket 连接成功后，打印日志
  socket.onclose = function (event) {
    console.log("Client notified socket has closed", event);
  };
  // 关闭 WebSocket 连接失败后，打印日志
  socket.onerror = function (event) {
    console.log("Client notified socket has errored", event);
  };
}
</script>
