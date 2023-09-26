import {defineStore} from "pinia";

export const useStore = defineStore('gvb', {
    state: () => {
        return {
            theme: true,
        }
    },
    actions: {
        setTheme() {
            this.theme = !this.theme
            if (this.theme) {
                document.documentElement.classList.remove("dark")
                localStorage.setItem("theme", "light")
                return
            }
            document.documentElement.classList.add("dark")
            localStorage.setItem("theme", "dark")
        },
        loadTheme() {
            const theme = localStorage.getItem("theme")
            if(theme === 'dark'){
                this.theme = false
                document.documentElement.classList.add("dark")
                return
            }
            this.theme = true
            document.documentElement.classList.remove("dark")
        }
    }


})