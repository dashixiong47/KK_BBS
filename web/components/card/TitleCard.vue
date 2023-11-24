<template>
  <div
    class="glass text-light-5 w-60 flex flex-col rounded-2xl px-5 py-3 dark:text-light-1"
  >
    <div class="p-2 flex-shrink-0 border-b">{{ title }}</div>
    <ul class="w-full h-full py-2 overflow-y-auto hidden-scrollbar">
      <li
        v-for="item in list"
        @click="change(item)"
        class="cursor-pointer hover:bg-blue-100 rounded-2xl h-12 flex items-center px-2"
        :class="modelValue === item.id ? 'bg-blue-100' : ''"
      >
        <KLink to="#" class="my-4 w-full flex items-center">
          <div
            class="flex-shrink-0 w-10 h-10 flex items-center justify-center rounded-full bg-gradient-to-t from-indigo-200 via-purple-200 to-pink-200"
          >
            <img v-if="item.icon" src="" alt="" srcset="" />
            <div v-else>{{ $t("default") }}</div>
          </div>
          <span
            class="ml-3 text-md w-full truncate whitespace-nowrap hover:text-blue-400"
          >
            {{ item.name }}-{{ item.id }}
          </span>
        </KLink>
      </li>
    </ul>
  </div>
</template>

<script setup>
let { list } = defineProps({
  title: {
    type: String,
    default: "More Pages",
  },
  list: {
    type: Array,
    default: () => [],
  },
  modelValue: {
    type: String,
    default: 0,
  },
});
let emit = defineEmits(["update:modelValue", "change"]);
const change = (data) => {
  emit("update:modelValue", data.id);
  emit("change", data.id);
};
</script>
