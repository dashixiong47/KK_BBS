<template>
  <div>
    <p class="border-b pb-2 mb-2 font-semibold">排行榜</p>
    <ul>
      <li
        v-for="iten in 6"
        class="border-b flex items-center justify-between py-2"
      >
        <div>
          头像
          <!-- <div>name</div> -->
        </div>
        <button class="text-sm flex items-center border rounded-md px-5 py-1">
          关注
        </button>
      </li>
    </ul>
  </div>
</template>

<script setup></script>
