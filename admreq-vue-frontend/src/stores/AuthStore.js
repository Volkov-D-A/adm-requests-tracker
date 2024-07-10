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
            if (date.getUTCDate() < 10) {
                result = "0" + date.getUTCDate();
            } else {
                result = date.getUTCDate();
            }
            result += "."
            if (date.getUTCMonth() < 10) {
                result = result + "0" + (date.getUTCMonth()+1);
            } else {
                result = result + (date.getUTCMonth()+1);
            }
            result += "." + date.getUTCFullYear() + " ";
            if (date.getUTCHours() < 10) {
                result = result + "0" + date.getUTCHours();
            } else {
                result = result + date.getUTCHours();
            }
            result += ":"
            if (date.getUTCMinutes() < 10) {
                result = result + "0" + date.getUTCMinutes();
            } else {
                result = result + date.getUTCMinutes();
            }
            return result
        },
    }
})