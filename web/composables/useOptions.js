// composables/useOptions.js
import { ref, computed } from 'vue';

export default function useOptions() {
    let typeOptions = [
        {
            name: "全部",
            type: "all",
            value: -1,
        },
        {
            name: "帖子",
            value: 1,
            type: "topic",
        },
        {
            name: "图片",
            type: "img",
            value: 2,
        },
        {
            name: "视频",
            type: "video",
            value: 3,
        },
        {
            name: "小说",
            type: "text",
            value: 4,
        },
    ]
    return {
        typeOptions
    }
}