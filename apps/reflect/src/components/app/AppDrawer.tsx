import { styled, useTheme } from "@mui/material/styles";
import {
  Drawer,
  IconButton,
  Divider,
  Typography,
  MenuList,
  MenuItem,
  ListItemIcon,
  ListItemText,
  Stack,
} from "@mui/material";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import { env } from "../../env/client.mjs";
import { RiGlobalFill } from "react-icons/ri";
import Link from "next/link";

const DrawerHeader = styled("div")(({ theme }) => ({
  display: "flex",
  alignItems: "center",
  padding: theme.spacing(0, 1),
  ...theme.mixins.toolbar,
  justifyContent: "flex-end",
}));

interface IAppDrawer {
  isDrawerOpen: boolean;
  drawerHandler: () => void;
}

const AppDrawer = ({ isDrawerOpen, drawerHandler }: IAppDrawer) => {
  const theme = useTheme();

  return (
    <Drawer
      sx={{
        width: 240,
        flexShrink: 0,
        "& .MuiDrawer-paper": {
          width: 240,
          boxSizing: "border-box",
        },
      }}
      variant="persistent"
      anchor="left"
      open={isDrawerOpen}>
      <DrawerHeader>
        <Typography
          variant="h5"
          sx={{
            flexGrow: 1,
            ml: 2,
            fontFamily: "Source Code Pro",
            fontStyle: "italic",
            fontWeight: 700,
            letterSpacing: ".3rem",
            color: "inherit",
            textDecoration: "none",
          }}>
          {env.NEXT_PUBLIC_APP_NAME}
        </Typography>
        <IconButton onClick={drawerHandler}>
          {theme.direction === "ltr" ? (
            <ChevronLeftIcon />
          ) : (
            <ChevronRightIcon />
          )}
        </IconButton>
      </DrawerHeader>
      <Divider />
      <Stack direction="column" padding="20px 20px 0px 20px">
        <Typography fontSize="large">Common</Typography>
        <MenuList>
          <MenuItem href="/app/command-list" component={Link}>
            <ListItemIcon>
              <RiGlobalFill />
            </ListItemIcon>
            <ListItemText>All Command List</ListItemText>
          </MenuItem>
        </MenuList>
      </Stack>
    </Drawer>
  );
};

export default AppDrawer;
