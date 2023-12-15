<template>
  <div class="">
    <!-- 侧边状态栏 -->
    <Sidebar :detail="detail" />
    <div class="grid grid-cols-9 gap-5">
      <div class="relative h-full col-span-9 sm:col-span-6">
        <Card>
          <div class="flex" v-if="detail.user">
            <Avatar :url="detail.user?.avatar" />
            <div class="w-full ml-2 flex items-center justify-between">
              <div class="flex flex-col items-start">
                <KLink to="#" class="text-md font-bold font-main-color">
                  {{ detail.user?.nickname }}
                </KLink>
                <span class="text-xs font-regular-color">
                  发布时间:{{ getRelativeTime(detail.createdAt) }}
                </span>
              </div>
            </div>
          </div>
          <h2 class="font-medium text-xl mt-5">{{ detail.title }}</h2>
          <Texts
            v-if="detail.topicDetail.texts.length"
            :detail="detail.topicDetail"
            v-model="currentText"
          />
        </Card>
        <Card id="comment" class="mt-5">
          <Comments :topicId="route.params.id" @change="commentChange" />
        </Card>
      </div>
      <div class="sm:col-span-3 sm:block m-1">
        <div class="sticky top-0">
          <Card class="mb-5">
            <p class="border-b pb-2 mb-2 font-bold">文章列表</p>
            <ul>
              <li
                v-for="item in detail.topicDetail.texts"
                class="cursor-pointer truncate rounded-xl overflow-hidden px-2 hover:bg-[--illuminate-color] hover:text-white"
                :class="{
                  'text-[--illuminate-color]': currentText === item.order,
                }"
                @click="currentText = item.order"
              >
                {{ item.name }}
              </li>
            </ul>
          </Card>
          <Card>
            <p class="border-b pb-2 mb-2 font-bold">附件</p>
            <AttachmentList :topicId="route.params.id" />
          </Card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRoute } from "#app";
import { useGetTopicDetail } from "~/api/server";
import { useAppConfigStore } from "~/stores/init";
import Texts from "~/components/detail/Texts.vue";
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);

const { getRelativeTime } = useTime();
const route = useRoute();
let currentText = ref(0);
let detail = ref({
  topicDetail: {
    texts: [],
  },
});

async function getTopicDetail() {
  try {
    let { data } = await useGetTopicDetail(route.params.id);
    detail.value = data || {};
    useHead({
      title: detail.value.title + " - " + appConfig.value.appName,
    });
  } catch (error) {
    console.log(error);
    showError(error);
  }
}
function commentChange() {
  detail.value.comment = detail.value.comment + 1;
  detail.value.commentState = true;
  if (detail.value.topicDetail) {
    getTopicDetail();
  }
}
function init() {
  getTopicDetail();
}
init();
</script>
