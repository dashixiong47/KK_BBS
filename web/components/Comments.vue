<template>
  <div class="flex items-center mb-5">
    <div class="flex items-center">
      <span class="text-base font-extrabold">{{ $t("comments-title") }}</span>
      <span class="ml-2 text-sm font-normal text-light-5">32</span>
    </div>
    <div class="ml-5 text-sm text-light-5">
      <button
        class="border px-2 rounded-l-md"
        :class="{
          'bg-light-3 dark:bg-light-4': !selected,
          'text-blue-400': !selected,
        }"
        @click="change(0)"
      >
        {{ $t("comments-hot") }}
      </button>
      <button
        class="border px-2 rounded-r-md"
        :class="{
          'bg-light-3 dark:bg-light-4': selected,
          'text-blue-400': selected,
        }"
        @click="change(1)"
      >
        {{ $t("comments-new") }}
      </button>
    </div>
  </div>
  <ul>
    <li v-for="item in 10" class="flex border-b mb-5 pb-5">
      <Avatar />
      <div class="ml-5">
        <p class="mb-2"><KName /></p>
        <div class="text-sm">
          一、放下大概就是这样，即使我们没在一起，我也会好好的，谢谢时间惊艳了那段有你的记忆，也谢谢现在更努力变好的自己。
          二、抱歉啊，不能为你金戈铁马，也不能许你一世繁华，不过我能给你一个小家，里面温了杯暖茶。
          三、从晨昏到日暮，从清贫到富足，从少年到老迈，从相遇到余生，只想和你十指相扣，从此再不分开。
          四、你的名字，是我读过最短的情诗。我很喜欢你，像春去秋来，海棠花开。
          五、秒回的人应该很温柔吧，因为一直在等喜欢的人，也舍不得让喜欢的人等。
        </div>
        <div class="text-xs text-light-5">
          <span class="mr-2 cursor-pointer">8小时之前</span>
          <span class="mr-2 cursor-pointer">
            <Icon name="icon-dianzan" class="mr-1" />11
          </span>
          <span class="cursor-pointer">{{ $t("comments-reply") }}</span>
        </div>
        <div>
          <SubComments />
        </div>
      </div>
    </li>
  </ul>
</template>

<script setup>
let selected = ref(0);
const change = (num) => {
  selected.value = num;
};
</script>
