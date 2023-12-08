<template>
  <div class="pb-5">
    <ul class="m-1">
      <li
        v-for="item in list"
        class="shadow-center w-full p-5 rounded-2xl mb-5"
      >
        <Topic :detail="item" />
      </li>
    </ul>
  </div>
</template>

<script setup>
let props = defineProps({
  userId: {
    type: String,
    default: "",
  },
  type: {
    type: Number,
    default: "",
  },
  list: {
    type: Array,
    default: () => [],
  },
});
</script>
