<template>
  <nav class="h-16 w-full flex items-center px-2 justify-between border-b">
    <el-icon class="cursor-pointer" :size="20" @click="change">
      <Fold v-if="!isCollapse" />
      <Expand v-else />
    </el-icon>
    <el-avatar :size="50" :src="getUserInfo.avatar" />
  </nav>
</template>

<script setup>
import { computed, ref } from "vue";
import { useUserInfoStore } from "@/stores/userInfo";
import { userInfo } from "@/api/user.js";
const userInfoStore = useUserInfoStore();
const getUserInfo = computed(() => userInfoStore.getUserInfo);
const isCollapse = ref(false);
const emit = defineEmits(["switchCollapse"]);
function change() {
  isCollapse.value = !isCollapse.value;
  emit("switchCollapse", isCollapse.value);
}
async function getUser() {
  try {
    const { data } = await userInfo();
    userInfoStore.setUserInfo(data);
  } catch (error) {}
}
getUser();
</script>

<style lang="scss" scoped></style>
