import React, { useContext, useState } from "react";
import { Link, useHistory } from "react-router-dom";
import { ThemeContext } from "../theme/themeprovider";
import Button from "@material-ui/core/Button";
import * as ActionType from "../login/actionTypes";
import { AuthContext } from "../context/authContext";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import Badge from "@material-ui/core/Badge";
import MenuItem from "@material-ui/core/MenuItem";
import Menu from "@material-ui/core/Menu";
import MenuIcon from "@material-ui/icons/Menu";
import CssBaseline from "@material-ui/core/CssBaseline";
import CompanyIcon from "../general/reusable/icons/company";
import SearchIcon from "@material-ui/icons/Search";
import AccountCircle from "@material-ui/icons/AccountCircle";
import MailIcon from "@material-ui/icons/Mail";
import NotificationsIcon from "@material-ui/icons/Notifications";
import MoreIcon from "@material-ui/icons/MoreVert";
import { getThemeProps, createStyles } from "@material-ui/styles";
import useStyles from "./styles";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import { Theme, withStyles } from "@material-ui/core/styles";
import Switch, { SwitchClassKey, SwitchProps } from "@material-ui/core/Switch";

interface Styles extends Partial<Record<SwitchClassKey, string>> {
	focusVisible?: string;
}

interface Props extends SwitchProps {
	classes: Styles;
}

const IOSSwitch = withStyles((theme: Theme) =>
	createStyles({
		root: {
			width: 42,
			height: 26,
			padding: 0,
			margin: theme.spacing(1)
		},
		switchBase: {
			padding: 1,
			"&$checked": {
				transform: "translateX(16px)",
				color: theme.palette.common.white,
				"& + $track": {
					backgroundColor: "#52d869",
					opacity: 1,
					border: "none"
				}
			},
			"&$focusVisible $thumb": {
				color: "#52d869",
				border: "6px solid #fff"
			}
		},
		thumb: {
			width: 24,
			height: 24
		},
		track: {
			borderRadius: 26 / 2,
			border: `1px solid ${theme.palette.grey[400]}`,
			backgroundColor: theme.palette.grey[50],
			opacity: 1,
			transition: theme.transitions.create(["background-color", "border"])
		},
		checked: {},
		focusVisible: {}
	})
)(({ classes, ...props }: Props) => {
	return (
		<Switch
			focusVisibleClassName={classes.focusVisible}
			disableRipple
			classes={{
				root: classes.root,
				switchBase: classes.switchBase,
				thumb: classes.thumb,
				track: classes.track,
				checked: classes.checked
			}}
			{...props}
		/>
	);
});

const CustomNavbar: React.FC = () => {
	const { dispatch } = useContext(AuthContext);
	const setThemeName = useContext(ThemeContext);
	const [themeState, setThemestate] = useState({ checked: true });
	const handleUniversalThemeChange = (name: string) => (
		event: React.ChangeEvent<HTMLInputElement>
	) => {
		setThemestate({ ...themeState, [name]: event.target.checked });
		// setThemestate(!themeState)
		event.target.checked
			? setThemeName("lightTheme")
			: setThemeName("darkTheme");
	};
	const history = useHistory();

	const classes = useStyles("");
	const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
	const [
		mobileMoreAnchorEl,
		setMobileMoreAnchorEl
	] = React.useState<null | HTMLElement>(null);

	const isMenuOpen = Boolean(anchorEl);
	const isMobileMenuOpen = false;//Boolean(mobileMoreAnchorEl);

	const open = Boolean(anchorEl);

	const handleProfileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
		setAnchorEl(event.currentTarget);
	};

	const handleClose = () => {
		setAnchorEl(null);
	};

	const handleMobileMenuClose = () => {
		setMobileMoreAnchorEl(null);
	};

	const handleMenuClose = () => {
		setAnchorEl(null);
		handleMobileMenuClose();
	};

	const handleMobileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
		setMobileMoreAnchorEl(event.currentTarget);
	};

	const menuId = "primary-search-account-menu";
	const renderMenu = (
		<Menu
			anchorEl={anchorEl}
			anchorOrigin={{ vertical: "top", horizontal: "right" }}
			id={menuId}
			keepMounted
			transformOrigin={{ vertical: "top", horizontal: "right" }}
			open={isMenuOpen}
			onClose={handleMenuClose}
		>
			{/* <MenuItem onClick={handleMenuClose}>Profile</MenuItem> */}
			<MenuItem onClick={handleMenuClose}>My account</MenuItem>
		</Menu>
	);

	const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
		setAnchorEl(event.currentTarget);
	};

	const mobileMenuId = "primary-search-account-menu-mobile";
	const renderMobileMenu = (
		<Menu
			anchorEl={mobileMoreAnchorEl}
			anchorOrigin={{ vertical: "top", horizontal: "right" }}
			id={mobileMenuId}
			keepMounted
			transformOrigin={{ vertical: "top", horizontal: "right" }}
			open={isMobileMenuOpen}
			onClose={handleMobileMenuClose}
		>
			<MenuItem>
				<Link to="/about">
					<p>About</p>
				</Link>
			</MenuItem>
			<MenuItem>
				<Link to="/register">
					<p>Join</p>
				</Link>
			</MenuItem>
			<MenuItem>
				<Link to="/login">
					<p>login</p>
				</Link>
			</MenuItem>
		</Menu>
	);

	return (
		<>
			<CssBaseline />
			<div>
				<AppBar
					color="primary"
					elevation={1}
					className={classes.root}
				>
					<Toolbar variant="dense">
						<IconButton
							edge="start"
							className={classes.menuButton}
							color="primary"
							aria-label="open drawer"
						>
							{/* <MenuIcon /> */}
						</IconButton>
						<CompanyIcon />
						{localStorage.getItem("token") !== null && (
							<>
								<Link
									to="/discover"
									style={{
										textDecoration: "none",
										marginLeft: 10,
										color: "inherit"
									}}
								>
									<Button color="inherit">discover</Button>
								</Link>
								<Link
									to="/leaderboard"
									style={{
										textDecoration: "none",
										marginLeft: 10,
										color: "inherit"
									}}
								>
									<Button color="inherit">Leaderboard</Button>
								</Link>
							</>
						)}
						<div className={classes.search}>
							<div className={classes.searchIcon}>
								<SearchIcon />
							</div>
							<InputBase
								onChange={(event: any) => {
									//add stuff
									console.log(event);
								}}
								onSubmit={(event: any) => {
									event.preventDefault();
								}}
								placeholder="Search…"
								classes={{
									root: classes.inputRoot,
									input: classes.inputInput
								}}
								inputProps={{ "aria-label": "search" }}
							/>
						</div>
						<div className={classes.grow} />
						<div className={classes.sectionDesktop}>
							{/* <IconButton
								aria-label="show 4 new mails"
								color="secondary"
							>
								<Badge badgeContent={4} color="secondary">
									<MailIcon />
								</Badge>
							</IconButton> */}
							{!localStorage.getItem("token") && (
								<>
									<Link
										to="/register"
										style={{ textDecoration: "none" }}
									>
										<Button
											color="inherit"
											variant="outlined"
											style={{
												marginLeft: 10,
												color: "black"
											}}
										>
											Join
										</Button>
									</Link>
									<Link
										to="/login"
										style={{ textDecoration: "none" }}
									>
										<Button
											color="inherit"
											variant="contained"
											style={{
												marginLeft: 10,
												backgroundColor: "#5588FF"
											}}
										>
											login
										</Button>
									</Link>
									<Link
										to="/about"
										style={{
											textDecoration: "none",
											marginLeft: 10,
											color: "black"
										}}
									>
										<Button color="inherit">about</Button>
									</Link>
								</>
							)}
							{localStorage.getItem("token") !== null && (
								<>
									<IconButton
										aria-label="account of current user"
										aria-controls="menu-appbar"
										aria-haspopup="true"
										onClick={handleMenu}
										color="inherit"
									>
										<AccountCircle />
									</IconButton>
									<Menu
										id="menu-appbar"
										anchorEl={anchorEl}
										anchorOrigin={{
											vertical: "top",
											horizontal: "right"
										}}
										keepMounted
										transformOrigin={{
											vertical: "top",
											horizontal: "right"
										}}
										open={open}
										onClose={handleClose}
									>
											<Link
										to="/profile"
										style={{ textDecoration: "none" }}
									>
										<MenuItem onClick={handleClose}>
											Profile
										</MenuItem>
									</Link>
										{/*<MenuItem onClick={handleClose}>
											My account
										</MenuItem> */}
										{localStorage.getItem("token") !==
											null && (
											<>
												<Link
													to="/"
													style={{
														textDecoration: "none",
														marginLeft: 10,
														color: "black"
													}}
													onClick={() => {
														localStorage.removeItem(
															"token"
														);
														localStorage.removeItem(
															"isLoggedIn"
														);

														history.push("/");
													}}
												>
													<MenuItem color="inherit">
														logout
													</MenuItem>
												</Link>
												<FormControlLabel
													control={
														<IOSSwitch
															checked={
																themeState.checked
															}
															onChange={handleUniversalThemeChange(
																"checked"
															)}
															value="checked"
														/>
													}
													label={themeState.checked? "Light": "Dark"}
												/>
											</>
										)}
									</Menu>
								</>
							)}
						</div>
						<div className={classes.sectionMobile}>
							<IconButton
								aria-label="show more"
								aria-controls={mobileMenuId}
								aria-haspopup="true"
								onClick={handleMobileMenuOpen}
								color="secondary"
							>
								<MoreIcon />
							</IconButton>
						</div>
					</Toolbar>
				</AppBar>
				<Toolbar />
			</div>
			{renderMobileMenu}
			{/* {renderMenu} */}
		</>
	);
};

export default CustomNavbar;
