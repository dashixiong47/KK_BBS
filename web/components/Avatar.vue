<script setup>
const { getPath } = usePath();
let props = defineProps({
  size: {
    type: Number,
    default: 1,
  },
  url: {
    type: String,
    default: "",
  },
  userInfo: {
    type: [Object, Boolean],
    default: false,
  },
});
const sizeClass = computed(() => {
  switch (props.size) {
    case 1:
      return `w-11 h-11`;

    default:
      return `w-10 h-10`;
  }
});
const getRouterPath = computed(() => {
  return `/user/${props.userInfo.id}`;
});
</script>

<template>
  <div
    class="shadow-center rounded-full overflow-hidden bg-slate-200 flex-shrink-0"
    :class="sizeClass"
  >
    <Tippy v-if="userInfo" :arrow="false">
      <img
        :src="getPath(props.url)"
        alt=""
        srcset=""
        class="w-full h-full object-cover block"
      />

      <template #content>
        <div
          class="w-52 flex flex-col justify-center items-center shadow-center rounded-xl p-5"
        >
          <KLink :to="getRouterPath">
            <img
              :src="getPath(userInfo.avatar)"
              class="w-16 h-16 rounded-full"
              alt=""
              srcset=""
            />
          </KLink>
          <span class="text-sm mt-2">
            金币:
            <span class="text-red-500">
              {{ userInfo.coins }}
              <Icon name="healthicons:coins-outline"></Icon>
            </span>
          </span>
          <p class="mt-2 truncate w-full">{{ userInfo.introduction }}</p>
          <div class="w-full flex justify-between mt-2">
            <KLink :to="getRouterPath" class="flex flex-col items-center">
              <span class="font-bold">{{ userInfo.topicTotal }}</span>
              <span class="text-xs">动态</span>
            </KLink>
            <KLink
              :to="getRouterPath + '/follow'"
              class="flex flex-col items-center"
            >
              <span class="font-bold">{{ userInfo.follow }}</span>
              <span class="text-xs">关注</span>
            </KLink>
            <KLink
              :to="getRouterPath + '/fans'"
              class="flex flex-col items-center"
            >
              <span class="font-bold">{{ userInfo.fans }}</span>
              <span class="text-xs">粉丝</span>
            </KLink>
          </div>
        </div>
      </template>
    </Tippy>
    <img v-else :src="getPath(props.url)" alt="" srcset="" />
  </div>
</template>
