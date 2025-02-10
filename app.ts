import { Bot } from "grammy";
import * as process from "node:process";

const KEY= process.env.BOT_KEY || "";

const bot = new Bot(KEY);

bot.command("start", (ctx) => ctx.reply("Welcome! Up and running."));

bot.on("message", (ctx) => ctx.reply("Got another message!"));

bot.start();