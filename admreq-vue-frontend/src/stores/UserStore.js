import { defineStore } from "pinia";

export const useUserStore = defineStore('UserStore', {
    state: () => ({
        credentials: [],
        authorized: false,
        userError: '',
    }),
    actions: {
        async getAuth(user, password) {
            const res = await fetch("http://192.168.141.62:8080/v1/user/auth",{
                method: "POST",
                body: JSON.stringify({
                    login:user, 
                    password:password
                })
            })
            console.log(res)            
            const data = await res.json()
            if (res.status === 200) { 
                this.credentials = data
                console.log(this.credentials)
                this.authorized = true
            } else {
                this.userError = data.message
            }
        }
    }
})