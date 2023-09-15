<template>
  <div class="">
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9 sm:col-span-6">
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
              <div>
                <span class="mr-5 regular-text text-xs">
                  <Icon name="tabler:eye" class="mr-1" />{{
                    formatNumber(11111)
                  }}
                </span>
                <span class="regular-text text-xs">
                  <Icon name="ic:round-location-on" class="mr-1" />上海
                </span>
              </div>
            </div>
          </div>
          <div
            class="border-t mt-5 py-5"
            v-html="detail.topicDetail?.content"
          ></div>
          <div v-if="detail.topicDetail?.hidden">
            <div class="text-md font-bold">隐藏内容</div>
            <div v-html="detail.topicDetail?.hidden_content"></div>
          </div>
        </Card>
        <Card class="mt-5">
          <Comments />
        </Card>
      </div>
      <div class="hidden sm:col-span-3 sm:block m-1">
        <Card class="sticky top-0">
          <p class="border-b pb-2 mb-2 font-bold">附件</p>
          <ul>
            <li>121212</li>
          </ul>
        </Card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRoute } from "#app";
import { useGetTopicDetail } from "~/api/server";
import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
const { formatNumber } = useFormatNumber();
const { isMobile } = useMobileDetect();
const { getRelativeTime } = useTime();
const route = useRoute();
let detail = ref({});
definePageMeta({
  layout: "user",
});
async function init() {
  try {
    // const {data} = await useFetch("http://localhost:8080/api/v1/topic/" + route.params.id);
    let { data } = await useGetTopicDetail(route.params.id);
    detail.value = data || {};
    console.log(data);
  } catch (error) {
    console.log(error);
  }
}
init();
</script>
