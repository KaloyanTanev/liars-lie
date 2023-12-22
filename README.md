# Liars lie

Liars lie is a game where a player tries to guess the number from a set of agents, some of them are liars, some of them are honest.

## CLI Commands

### start --value v --max-value max --num-agents number --liar-ratio ratio

Launch agents and write their endpoints in `agents.config` file. Currently agents are hardcoded to start on ports :13000 and above.

#### value

Positive integer.
The truthful network value chosen by honest agents.

#### max-value

Positive integer.
Maximum value a lying agent can choose. A lying agent does not change its lying value during `play` round once it is started.

#### num-agents

Positive integer.
Number of agents.

#### liar-ratio

Positive float in the range of 0 <= liar-ratio <= 1.
Ratio of liars in the lying agents from total set of agents (specified in num-agents).

### play

Play one round of the game, outputiing a list of the values received by each agent and the concluded resulting value.

### stop

Stop all agents and delete `agents.config`.

### extend --value v --max-value max --num-agents number --liar-ratio ratio

Extend the set of agents in the game.

#### value

Positive integer.
The truthful network value chosen by honest agents.

#### max-value

Positive integer.
Maximum value a lying agent can choose. A lying agent does not change its lying value during `play` round once it is started.

#### num-agents

Positive integer.
Number of new agents.

#### liar-ratio

Positive float in the range of 0 <= liar-ratio <= 1.
Ratio of liars in the lying agents from the new set of agents (specified in num-agents).

### play-expert --num-agents number --liar-ratio ratio

Play in expert mode. In this mode agents communicate with themselves to synchronize and choose one exact value between themselves.

#### num-agents

Positive integer in the range of 0 <= num-agents <= total number of agents
Player chooses with how many of the agents it wants to play. They are chosen randomly, thus the ratio of honest and lying is random.

#### liar-ratio

Positive float in the range of 0 <= liar-ratio <= 1.
Expected ratio from the player of lying agents in the subset chosen. It helps to predict which is the truthful value from the set of responses.

### kill --id id

Kill a single agent.

#### id

Id of an agent.

## Build

Build the game. Required golang golang >=1.2.1.

```sh
scripts/build.sh
```

## Run

### From script as one command

Run the game specifying multiple commands in one line. Include one argument, separating each command with `|`.

Example usage:

```sh
scripts/run.sh 'start --value=42 --max-value=100 --num-agents=10 --liar-ratio=0.4 | play | stop'
```

```sh
scripts/run.sh 'start --value=20 --max-value=100 --num-agents=10 --liar-ratio=0.4 | play | extend --value=42 --max-value=100 --num-agents=5 --liar-ratio=0.2 | play | playexpert --num-agents=10 --liar-ratio=0.4 | play | kill --id=4 | playexpert --num-agents=10 --liar-ratio=0.4 | stop'
```

### From executable

Alternatively, you can run them one by one by executing `app/app`.

Example usage:

```sh
./app/app start --value=42 --max-value=100 --num-agents=100 --liar-ratio=0.2 & # wait for `ready` signal
./app/app playexpert --num-agents=40 --liar-ratio=0.6
./app/app playexpert --num-agents=50 --liar-ratio=0.6
./app/app playexpert --num-agents=80 --liar-ratio=0.3
./app/app kill --id=42
./app/app playexpert --num-agents=100 --liar-ratio=0.2
./app/app stop
```

## Test

```sh
scripts/unit_test.sh
```
