package starbot.commands;

import net.dv8tion.jda.core.events.message.guild.GuildMessageReceivedEvent;
import starbot.resources.interfaces.Command;
import starbot.util.Messages;

public class SexyGandalf implements Command {

    private static SexyGandalf instance;

    public static synchronized SexyGandalf instance() {
        if (instance == null)
            instance = new SexyGandalf();
        return instance;
    }

    private SexyGandalf() {
    }

    @Override
    public boolean call(String cmd, String[] args, GuildMessageReceivedEvent evt) {
        System.out.println("Called Sexy Gandalf!\n");
        return true;
    }

    @Override
    public boolean execute(String[] args, GuildMessageReceivedEvent evt) {
        evt.getMessage().getChannel().sendMessage(Messages.COMMAND_SEXYGANDALF).queue();
        return true;
    }

    @Override
    public void finish(boolean permission, boolean success, GuildMessageReceivedEvent evt) {
        System.out.println("Executed SexyGandalf:");
        System.out.println("Permission: " + permission);
        System.out.println("Success: " + success);
        System.out.println("\n");
    }

}