import { View, Text, TextInput, Button } from "react-native"
import * as yup from "yup"
import { yupResolver } from "@hookform/resolvers/yup"
import { useForm } from "react-hook-form"
import Input from "./input"
const schema = yup
  .object({
    username: yup.string().required('This field is required'),
    password: yup.string().required('This field is required')
  })
  .required()

const Login = () => {
    const { control, handleSubmit } = useForm({resolver: yupResolver(schema)});

    const onSubmit = (data) => {
        const { username, password } = data;
    }
    return (    
        <View style={{ paddingLeft: 16, paddingRight: 16, borderColor: 'green', borderWidth: 3, height: '100%' }}>
            <Input label='Username' name='username' control={control} />
            <Input label='Password' name='password' control={control} />
            <View style={{ marginTop: 8,}}>
                <Button title="Login" type="submit" onPress={handleSubmit(onSubmit)} />
            </View>
        </View>
    )
}

export default Login