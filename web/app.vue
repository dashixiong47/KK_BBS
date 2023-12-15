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
import { useUserStore } from "~/stores/main";
import { useAppConfigStore } from "~/stores/init";
const { initSocket, closeSocket } = useWebSocket();
let app = useAppConfigStore();
app.fetchHost();
let userStore = useUserStore();

let isLogin = computed(() => userStore.isLogin);

watch(
  () => isLogin.value,
  (val) => {
    if (val) {
      initSocket();
    } else {
      closeSocket();
    }
  },
  { immediate: true }
);
</script>
