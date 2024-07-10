import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useUsersStore = defineStore('UsersStore', {
    state: () => ({
        users: [],
        usersErrors: "",
    }),
    actions:{
        async getUsers(token) {
            const res = await fetch(path+'users/'+token,{
                method: "GET",
            })
            const data = await res.json()
            if (res.status === 200) {
                this.users = data.users
            }
        },
        async createUser(fn, ln, sn, dp, login, pass, role, token) {
            const res = await fetch(path+'user', {
                method: "POST",
                body: JSON.stringify({
                    firstname: fn,
                    lastname: ln,
                    surname: sn,
                    department: dp,
                    login: login,
                    password: pass,
                    role: role,
                    token: token
                })
            })
            if(res.status === 200) {
                this.getUsers(token)    
            } else {
                const data = await res.json()
                this.usersErrors = data.message
            }
        },
        getEmployeeItems() {
            console.log("all users:", this.users)
            var y = []
            const employ = this.users.filter((el) => el.Role != "user")
            for (let i = 0; i < employ.length; i++) {
                const x = {
                    title: employ[i].lastname + " " + employ[i].firstname[0] + "." + employ[i].surname[0] + ".",
                    value: employ[i].uuid
                }
                y.push(x)
           }
           console.log("employes:", y)
           return y
        },
        async deleteUser(userid, token) {
            console.log(userid)
            const res = await fetch(path+'userdel',{
                method: "POST",
                body: JSON.stringify({
                    uuid: userid,
                    token: token,
                })
            })
            if (res.status === 200) {
                this.getUsers(token)
            }
        }
    }, 
})