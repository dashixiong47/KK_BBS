import { KLink } from '.nuxt/components';
<template>
  <div class="border-t mt-5 pt-5">
    <ul class="flex">
      <li v-for="item in types">
        <KLink
          :to="`/${item.type}/create`"
          class="mr-5"
        >
          <KButton class="mr-5">
            {{ item.name }}
          </KButton>
        </KLink>
      </li>
    </ul>
  </div>
</template>

<script setup>
const types = [
  {
    name: "视频",
    type: "video",
  },
  {
    name: "文章",
    type: "text",
  },
];
</script>
