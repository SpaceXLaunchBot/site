import { createMuiTheme } from '@material-ui/core/styles';

// Based on https://discord.com/branding
const theme = createMuiTheme({
    palette: {
        primary: { main: '#23272A' },
        secondary: { main: '#7289DA' },
        text: { primary: '#FFFFFF', secondary: '#686d73' },
        background: { default: '#2C2F33' },
    },
});

export default theme;
