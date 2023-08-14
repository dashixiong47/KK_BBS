<template>
  <ul class="mt-2">
    <li v-for="item in 1" class="flex mb-2 pb-2">
      <Avatar class="w-6 h-6" />
      <div class="ml-2">
        <div class="text-sm">
          <span class="mb-2">
            <Name />
            <span class="flex-shrink-0 mx-2">
              {{ $t("subComments-reply") }}
            </span>
            <Name name="@name" class="text-blue-400" /> :
          </span>
          一、放下大概就是这样，即使我们没在一起，我也会好好的，谢谢时间惊艳了那段有你的记忆，也谢谢现在更努力变好的自己。
        </div>
        <div class="text-xs text-light-5">
          <span class="mr-2 cursor-pointer">8小时之前</span>
          <span class="mr-2 cursor-pointer">
            <Icon name="icon-dianzan" class="mr-1" />11
          </span>
          <span class="cursor-pointer">{{ $t("subComments-reply") }}</span>
        </div>
      </div>
    </li>
  </ul>
</template>

<script setup></script>
