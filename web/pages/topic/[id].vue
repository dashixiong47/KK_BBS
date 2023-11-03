<template>
  <div class="">
    <!-- 侧边状态栏 -->
    <Sidebar :detail="detail" />
    <div class="grid grid-cols-9">
      <div class="relative h-full m-1 col-span-9 sm:col-span-6">
        <Card>
          <div class="flex" v-if="detail.user">
            <Avatar :url="detail.user?.avatar" />
            <div class="w-full ml-2 flex items-center justify-between">
              <div class="flex flex-col items-start">
                <KLink to="#" class="text-md font-bold primary-text">
                  {{ detail.user?.nickname }}
                </KLink>
                <span class="text-xs secondary-text">
                  发布时间:{{ getRelativeTime(detail.createdAt) }}
                </span>
              </div>
              <!-- <div>
                <span class="mr-5 regular-text text-xs">
                  <Icon name="tabler:eye" class="mr-1" />{{
                    formatNumber(detail.view)
                  }}
                </span>
                <span class="regular-text text-xs">
                  <Icon name="ic:round-location-on" class="mr-1" />上海
                </span>
              </div> -->
            </div>
          </div>
          <h2 class="font-medium text-xl mt-5">{{ detail.title }}</h2>
          <div class="py-5" v-html="detail.topicDetail?.content"></div>
          <div v-if="detail.topicDetail?.hidden">
            <div class="text-md font-bold">隐藏内容</div>
            <div v-html="detail.topicDetail?.hidden_content"></div>
          </div>
        </Card>
        <Card id="comment" class="mt-5">
          <Comments :topicId="route.params.id" @change="commentChange" />
        </Card>
      </div>
      <div class="sm:col-span-3 sm:block m-1">
        <div class="sticky top-0">
          <Card>
            <p class="border-b pb-2 mb-2 font-bold">附件</p>
            <ul>
              <li>121212</li>
            </ul>
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

import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);
definePageMeta({
  layout: "user",
});
const { formatNumber } = useFormatNumber();
const { isMobile } = useMobileDetect();
const { getRelativeTime } = useTime();
const route = useRoute();

let detail = ref({});
async function getTopicDetail() {
  try {
    // const {data} = await useFetch("http://localhost:8080/api/v1/topic/" + route.params.id);
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
}
function init() {
  getTopicDetail();
}
init();
</script>
