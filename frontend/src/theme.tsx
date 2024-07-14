import { extendTheme } from "@chakra-ui/react";


export const theme = extendTheme({
    styles: {
        global: {
            'html, body': {
                color: 'gray.600',
                lineHeight: 'tall',
            },
            a: {
                color: 'teal.500',
            },
        },
    },
});
