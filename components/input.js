import { useForm , useController} from "react-hook-form"
import { View, Text, TextInput } from "react-native"
const Input = ({ name, control, label}) => {
    const { field, fieldState } = useController({
        control,
        defaultValue: '',
        name,
    })
    return (
        <View>
            <Text style={{ marginTop: 4, fontSize: 16}}>{label}</Text>
            <TextInput
                value={field.value}
                onChangeText={field.onChange}
                style={{borderWidth: 1, borderColor: 'black', marginTop: 2}}
                name={field.name}
            />
            {fieldState.error && fieldState.isDirty && 
                <Text style={{ color: 'red', fontSize: 16 }}>{fieldState.error.message}</Text>}
         </View>
    )
}

export default Input