<template>
    <v-row>
        <v-col cols="6">
            <div class="pa-3">
                <v-list lines="one">
                    <v-list-item 
                        v-for="user in UsersStore.users"
                        :key="user.uuid"
                    >
                        <template v-slot:append>
                            <!-- <v-icon  color="red" icon="mdi-comment-alert" @click="UsersStore.deleteUser(user.uuid, AuthStore.credentials.token)"></v-icon> -->
                            <v-icon  color="red" icon="mdi-comment-alert" @click="delid = user.uuid, dialog = true"></v-icon>
                        </template>
                        <v-list-item-title>
                            {{ user.lastname }} {{ user.firstname }} {{ user.surname }}
                        </v-list-item-title>
                        <v-list-item-subtitle>
                            Роль: {{ user.Role }} / Подразделение: {{ user.departmentName}}
                        </v-list-item-subtitle>

                    </v-list-item>
                </v-list>
                <v-dialog v-model="dialog" width="500">
                    <v-card
                        prepend-icon="mdi-trash-can"
                        text="Вы уверены, что хотите заблокировать пользователя? Доступ к системе обращений будет прекращен!"
                        title="Блокировка пользователя"
                    >
                    <template v-slot:actions>
                    <v-spacer></v-spacer>

                <v-btn fab dark small color="primary" @click="dialog = false">
                Отменить
                </v-btn>

                <v-btn fab dark small color="red" @click="UsersStore.deleteUser(delid, AuthStore.credentials.token), dialog = false">
                    Удалить
                </v-btn>
            </template>
            </v-card>
        </v-dialog>
            </div>
        </v-col>
        <v-col cols="3">
            <v-card class="pa-3 ma-3" variant="elevated" elevation="16" color="teal-lighten-4">
                <span class="ma-2">Добавить отдел:</span><br><br>
            <v-form @submit.prevent="UsersStore.addDepartment(depart, dowork, AuthStore.credentials.token)">
                <v-text-field v-model="depart" label="Название отдела"></v-text-field>
                <v-switch label="Выполняет тикеты" v-model="dowork"></v-switch>
                <v-btn type="submit" color="teal-darken-1" block class="mt-2">Добавить</v-btn>
            </v-form>
            </v-card>
        </v-col>
        <v-col cols="3">
                <v-card class="pa-3 ma-3" variant="elevated" elevation="16" color="teal-lighten-4">
                    <span class="ma-2">Добавить пользователя:</span><br><br>
                <v-form fast-fail @submit.prevent="UsersStore.createUser(firstname, lastname, surname, department, username, password, role, AuthStore.credentials.token)">
                    <v-text-field v-model="lastname" label="Фамилия"></v-text-field>
                    <v-text-field v-model="firstname" label="Имя"></v-text-field>
                    <v-text-field v-model="surname" label="Отчество"></v-text-field>
                    <v-select
                        label="Отдел"
                        v-model="department"
                        :items="UsersStore.departments"
                        item-title="department"
                        item-value="uuid"
                    >
                    </v-select>
                    <v-text-field v-model="username" label="Логин"></v-text-field>
                    <v-text-field v-model="password" label="Пароль"></v-text-field>
                    <v-select
                        label="Роль"
                        v-model="role"
                        :items="[{title: 'Пользователь', value: 'user'}, {title: 'Исполнитель', value: 'employee'}, {title: 'Администратор', value: 'admin'},]"
                    ></v-select>
                    <v-btn type="submit" color="teal-darken-1" block class="mt-2">Добавить</v-btn>
                </v-form>
                <span class="d-flex align-center justify-center ma-3 text-red">{{ UsersStore.usersErrors }}</span>
                </v-card>
        </v-col>
    </v-row>
</template>

<script setup>
import { useAuthStore } from '../stores/AuthStore';
import { useUsersStore } from '../stores/UsersStore';
import { ref } from 'vue';

const AuthStore = useAuthStore();
const UsersStore = useUsersStore();

UsersStore.getUsers(AuthStore.credentials.token);
UsersStore.getDepartments(AuthStore.credentials.token)

const firstname = ref("");
const lastname = ref("");
const surname = ref("");
const department = ref("");
const username = ref("");
const password = ref("");
const role = ref("");
const dowork = ref(false);
const depart = ref("");
const dialog = ref(false)
var delid = ""
</script>