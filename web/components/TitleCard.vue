<template>
  <div class="bg-white w-60 shadow-center rounded-2xl px-5 py-3 text-gray-400">
    <div class="">More Pages</div>
    <ul class="w-full h-full">
      <li v-for="item in list">
        <NuxtLink to="#" class="my-4 w-full flex items-center">
          <div
            class="flex-shrink-0 w-12 h-12 flex items-center justify-center rounded-full bg-gradient-to-t from-indigo-200 via-purple-200 to-pink-200"
          >
            <Icon class="text-white" name="icon-xiaoxi" size="text-4xl mx-1" />
          </div>
          <span class="ml-3 text-md w-full truncate whitespace-nowrap"
            >1212</span
          >
        </NuxtLink>
      </li>
    </ul>
  </div>
</template>

<script setup>
const list = [0, 1, 2, 3, 4, 0, 1, 2, 3, 4];
</script>
