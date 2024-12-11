import { View, Text, TextInput, Button } from "react-native"
import * as yup from "yup"
import { yupResolver } from "@hookform/resolvers/yup"
import { useForm } from "react-hook-form"
import Input from "./input"
import axios from 'axios'

const schema = yup
  .object({
    //email: yup.string().required('Email is required'),
    username: yup.string().min(3, 'Username min length is 3 letters').required('Username is required'),
    password: yup.string().min(6, 'Password min length is 6 letters').required('Password is required'),
    password2: yup.string().min(6, 'Password min length is 6 letters').required('Password is required')
  })
  .required()

const Register = () => {
    const { control, handleSubmit } = useForm({resolver: yupResolver(schema)});

    const onSubmit = (data) => {
        const { username, password, password2 } = data;
    }
    return (
        <View style={{ paddingLeft: 16, paddingRight: 16, borderColor: 'green', borderWidth: 3, height: '100%' }}>
            <Input label='Username' name='username' control={control} />
            <Input label='Password' name='password' control={control} />
            <Input label='Password2' name='password2' control={control} />
            <View style={{ marginTop: 8,}}>
                <Button title="Login" type="submit" onPress={handleSubmit(onSubmit)} />
            </View>
        </View>
    )
}

export default Register