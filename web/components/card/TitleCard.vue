<template>
  <div
    class="bg-light-1 text-light-5 w-60 flex flex-col shadow-center rounded-2xl px-5 py-3 dark:bg-dark-2 dark:text-light-1"
  >
    <div class="p-2 flex-shrink-0 border-b">More Pages</div>
    <ul class="w-full h-full overflow-y-auto hidden-scrollbar">
      <li v-for="item in list" class="hover:bg-blue-100 rounded-2xl h-12 flex items-center px-2">
        <KLink to="#" class="my-4 w-full flex items-center">
          <div
            class="flex-shrink-0 w-10 h-10 flex items-center justify-center rounded-full bg-gradient-to-t from-indigo-200 via-purple-200 to-pink-200"
          >
            <Icon class="text-white" name="icon-xiaoxi" size="text-4xl mx-1" />
          </div>
          <span class="ml-3 text-md w-full truncate whitespace-nowrap hover:text-blue-400"
            >1212</span
          >
        </KLink>
      </li>
    </ul>
  </div>
</template>

<script setup>
defineProps({});
const list = [0, 1, 2, 3, 4, 0, 1, 2, 3, 4];
</script>
