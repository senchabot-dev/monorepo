import { IconButton, Typography } from "@mui/material";
import { SiTwitch } from "react-icons/si";
import { BootstrapTooltip } from "src/components/Tooltip";
import { trpc } from "../../../../utils/trpc";

const GetTwitchBotButton = () => {
  const { data: twitchAcc } = trpc.check.checkTwitchAcc.useQuery();

  const twitchBotMutate = trpc.twitchBot.add.useMutation({
    onSuccess() {
      alert("Twitch bot added");
    },

    onError(error: any) {
      if (!error.shape) return;
      alert(error.shape.message);
    },
  });

  return (
    <BootstrapTooltip title="Get Twitch Bot">
      <Typography
        onClick={() =>
          !twitchAcc
            ? alert(
                "Before you can add the Twitch bot, you need to link your Twitch account in Settings/Security section.",
              )
            : twitchBotMutate.mutate()
        }>
        <IconButton
          aria-label="get twitch bot"
          sx={{
            display: "flex",
          }}>
          <SiTwitch />
        </IconButton>
      </Typography>
    </BootstrapTooltip>
  );
};

export default GetTwitchBotButton;
