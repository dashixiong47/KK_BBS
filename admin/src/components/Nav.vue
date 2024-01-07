<template>
  <nav class="h-16 w-full flex items-center px-2 justify-between border-b">
    <el-icon class="cursor-pointer" :size="20" @click="change">
      <Fold v-if="!isCollapse" />
      <Expand v-else />
    </el-icon>

    <el-dropdown @command="command">
      <!-- {{ getPath(getUserInfo.avatar) }} -->
      <span class="el-dropdown-link">
        <el-avatar :size="50" :src="getPath(getUserInfo.avatar)">
          <el-icon :size="30"><User /></el-icon>
        </el-avatar>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item command="quit">退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </nav>
</template>

<script setup>
import { computed, ref } from "vue";
import { useUserInfoStore } from "@/stores/userInfo";
import { userInfo } from "@/api/user.js";
import { getPath } from "@/utils/path.js";
import { useRouter } from "vue-router";
const router = useRouter();
const userInfoStore = useUserInfoStore();
const getUserInfo = computed(() => userInfoStore.getUserInfo);
const isCollapse = ref(false);
const emit = defineEmits(["switchCollapse"]);
function change() {
  isCollapse.value = !isCollapse.value;
  emit("switchCollapse", isCollapse.value);
}
function command(type) {
  switch (type) {
    case "quit":
      quit();
      break;
    default:
      break;
  }
}
function quit() {
  localStorage.removeItem("token");
  router.replace({ path: "/login" });
}
async function getUser() {
  try {
    const { data } = await userInfo();
    console.log(data,"data");
    userInfoStore.setUserInfo(data);
  } catch (error) {}
}
getUser();
</script>

<style lang="scss" scoped></style>
