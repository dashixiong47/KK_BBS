<template>
  <ul class="grid grid-cols-12 gap-4">
    <li
      v-for="item in props.list"
      class="relative col-span-4 sm:col-span-3 lg:col-span-3"
    >
      <KLink :to="`/text/${item.id}`">
        <div
          class="shadow-center flex items-center justify-center h-40 sm:h-60 rounded-xl relative overflow-hidden"
        >
          <KImage
            class="absolute z-0 left-0 top-0 w-full h-full"
            v-for="val in item.covers"
            :source="getPath(val)"
          />
          <span v-if="!item.covers.length" class="text-[--font-secondary-color] dark:text-[--dark-font-secondary-color]">暂无封面</span>
        </div>
        <div class="truncate text-sm mt-2">{{ item.title }}</div>
        <div
          class="flex truncate text-xs text-[--font-secondary-color] dark:text-[--dark-font-secondary-color]"
        >
          <span class="block truncate w-[70%]">{{ item.user.nickname }}</span>
          <span class="block w-[30%]">{{ getRelativeTime(item.createdAt) }}</span>
        </div>
      </KLink>
    </li>
  </ul>
</template>

<script setup>
const { getPath } = usePath();
const { getRelativeTime } = useTime();
const props = defineProps({
  list: {
    type: Array,
    default: [],
  },
});
</script>
