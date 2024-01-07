<script setup>
import { follow } from "~/api";
const { addMessage } = useMessage();
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
const followUser = async () => {
  try {
    await follow(props.userInfo.id);
    props.userInfo.isFollow = !props.userInfo.isFollow;
  } catch (error) {
    addMessage(error.message, "error");
  }
};
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

          <p class="mt-2 truncate w-full text-center">
            {{ userInfo.introduction }}
          </p>
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
          <div class="w-full flex items-center justify-between mt-2">
            <KButton
              size="mini"
              @click="followUser(userInfo.id)"
              class="flex items-center"
            >
              {{ userInfo.isFollow ? "已关注" : "关注" }}
              <Icon
                class="ml-1"
                :name="
                  userInfo.isFollow
                    ? 'ep:circle-check-filled'
                    : 'ep:circle-check'
                "
              ></Icon>
            </KButton>
            <KButton size="mini" class="flex items-center">
              私信
              <Icon class="ml-1" name="ep:chat-dot-round"></Icon>
            </KButton>
          </div>
        </div>
      </template>
    </Tippy>
    <img v-else :src="getPath(props.url)" alt="" srcset="" />
  </div>
</template>
