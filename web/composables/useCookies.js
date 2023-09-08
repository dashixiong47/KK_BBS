// composables/useCookies.js

import { useCookie } from "#app"
export default function useCookies() {
    let setCookie = (key, value = null) => {
        console.log("setCookie", key, value);
        let cookie = useCookie(key);
        cookie.value = value
    }
    return { setCookie };
}
