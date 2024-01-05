<template>
    <div>
      <ul>
        <li
          v-for="item in list"
          class="flex justify-between shadow-center p-5 rounded-xl mb-5"
        >
          <div class="left flex">
            <div class="avatar mr-2">
              <Avatar
                :url="item.avatar"
                :userInfo="item"
                class="w-12 h-12 rounded-full"
              />
            </div>
            <div class="info">
              <div class="name font-extrabold">{{ item.nickname }}</div>
              <div class="desc text-sm">{{ item.introduction }}</div>
            </div>
          </div>
          <div class="right">
            <KButton @click="cancelFollow(item.id)">取消关注</KButton>
          </div>
        </li>
      </ul>
      <Pagination v-model="pages" @query="init" />
    </div>
  </template>
  
  <script setup>
  import { useAppConfigStore } from "~/stores/init";
  import { fansList } from "~/api/server";
  import { follow } from "~/api";
  const { addMessage } = useMessage();
  const route = useRoute();
  let appConfigStore = useAppConfigStore();
  let appConfig = computed(() => appConfigStore.getConfig);
  useHead({
    title: "关注 - " + appConfig.value.appName,
  });
  let list = ref([]);
  let pages = ref({
    total: 0,
    page: 1,
    pageSize: 10,
  });
  async function cancelFollow(id) {
    try {
      await follow(id);
      list.value = list.value.filter((item) => item.id !== id);
    } catch (error) {
      console.log(error);
      addMessage(error, "error");
    }
  }
  async function init() {
    try {
      let { data } = await fansList(route.params.id, pages);
      console.log(data);
      list.value = data.list;
      pages.value.total = data.total;
    } catch (error) {
      console.log(error);
    }
  }
  init();
  </script>
  