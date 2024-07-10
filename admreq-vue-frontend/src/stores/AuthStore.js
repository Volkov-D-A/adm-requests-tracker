import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useAuthStore = defineStore('AuthStore', {
    state: () => ({
        credentials: [],
        authorized: false,
        userError: '',
    }),
    actions: {
        async getAuth(user, password) {
            const res = await fetch(path+'user/auth',{
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
        },
        myDateTimeFormat(strdate) {
            var date = new Date(strdate)
            var result
            if (date.getDate() < 10) {
                result = "0" + date.getDate();
            } else {
                result = date.getDate();
            }
            result += "."
            if (date.getMonth() < 10) {
                result = result + "0" + date.getMonth();
            } else {
                result = result + date.getMonth();
            }
            result += "." + date.getFullYear() + " ";
            if (date.getHours() < 10) {
                result = result + "0" + date.getHours();
            } else {
                result = result + date.getHours();
            }
            result += ":"
            if (date.getMinutes() < 10) {
                result = result + "0" + date.getMinutes();
            } else {
                result = result + date.getMinutes();
            }
            return result
        },
    }
})