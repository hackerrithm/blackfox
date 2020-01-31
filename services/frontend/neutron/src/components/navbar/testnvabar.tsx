import React from "react";
import Box from "@material-ui/core/Box";
import CssBaseline from "@material-ui/core/CssBaseline";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";

export default function ButtonAppBar() {
        return (
                <div>
                        <CssBaseline />
                        <AppBar position="fixed">
                                <Toolbar>
                                        <Box mr={2}>
                                                <IconButton edge="start" color="inherit" aria-label="menu">
                                                        <MenuIcon />
                                                </IconButton>
                                        </Box>
                                        <Typography variant="h6">News</Typography>
                                        <Box flexGrow={1} />
                                        <Button color="inherit">Login</Button>
                                </Toolbar>
                        </AppBar>
                        <Toolbar />
    </div>
        );
}