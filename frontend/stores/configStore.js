import { defineStore } from "pinia";

export const useConfigStore = defineStore({
    id: "config-store",
    state: () => {
        return {
            mobile: false,
        }
    },
    actions: {
        disableMobile() {
            this.mobile = false
        },
        enableMobile() {
            this.mobile = true
        }
    },
    getters: {
        getMobile: state => state.mobile
    }
})