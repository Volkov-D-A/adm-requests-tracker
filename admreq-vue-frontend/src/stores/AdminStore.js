import { defineStore } from "pinia";

export const useAdminStore = defineStore('AdminStore', {
    state: () => ({
        users: [],
        usersErrors: "",
    }),
    actions:{
        async getUsers(token) {
            const res = await fetch("http://192.168.141.62:8080/v1/users/"+token,{
                method: "GET",
            })
            const data = await res.json()
            if (res.status === 200) {
                this.users = data.users
            }
        },
        async createUser(fn, ln, login, pass, role, token) {
            const res = await fetch("http://192.168.141.62:8080/v1/user", {
                method: "POST",
                body: JSON.stringify({
                    first_name: fn,
                    last_name: ln,
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
            var y = []
            const employ = this.users.filter((el) => el.Role != "user")
            for (let i = 0; i < employ.length; i++) {
                const x = {
                    title: employ[i].firstName + " " + employ[i].lastName,
                    value: employ[i].uuid
                }
                y.push(x)
           }
           return y
        }
    }, 
})