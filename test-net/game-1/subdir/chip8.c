#include "chip8.h"
#include <memory.h>
#include <assert.h>
#include <stdlib.h>
#include <time.h>
#include "SDL2/SDL.h"

const char chip8_default_character_set[] = {
    0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
    0x20, 0x60, 0x20, 0x20, 0x70, // 1
    0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
    0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
    0x90, 0x90, 0xF0, 0x10, 0x10, // 4
    0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
    0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
    0xF0, 0x10, 0x20, 0x40, 0x40, // 7
    0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
    0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
    0xF0, 0x90, 0xF0, 0x90, 0x90, // A
    0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
    0xF0, 0x80, 0x80, 0x80, 0xF0, // C
    0xE0, 0x90, 0x90, 0x90, 0xE0, // D
    0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
    0xF0, 0x80, 0xF0, 0x80, 0x80  // F
};

void chip8_init(struct chip8* chip8)
{
    memset(chip8, 0, sizeof(struct chip8));
    memcpy(&chip8->memory.memory, chip8_default_character_set, sizeof(chip8_default_character_set));
}

void chip8_load(struct chip8* chip8, const char* buf, size_t size)
{
    //Assert that we have enough memory
    assert(size+CHIP8_PROGRAM_LOAD_ADDRESS < CHIP8_MEMORY_SIZE);

    memcpy(&chip8->memory.memory[CHIP8_PROGRAM_LOAD_ADDRESS], buf, size);
    chip8->registers.PC = CHIP8_PROGRAM_LOAD_ADDRESS;
}

// nnn or addr - A 12-bit value, the lowest 12 bits of the instruction
// n or nibble - A 4-bit value, the lowest 4 bits of the instruction
// x - A 4-bit value, the lower 4 bits of the high byte of the instruction
// y - A 4-bit value, the upper 4 bits of the low byte of the instruction
// kk or byte - An 8-bit value, the lowest 8 bits of the instruction



void chip8_exec(struct chip8* chip8, unsigned short opcode)
{
    switch(opcode)
    {
        //Clear the display
        case 0x00E0:
            chip8_screen_clear(&chip8->screen);
            break;

        //Return from a subroutine
        case 0x00EE:
            chip8->registers.PC = chip8_stack_pop(chip8);
        break;

        default:
            chip8_exec_extended(chip8, opcode);

    }
}

void chip8_exec_extended(struct chip8* chip8, unsigned short opcode)
{

    unsigned short nnn = opcode & 0x0fff;
    unsigned char x = (opcode >> 8) & 0x000f;
    unsigned char y = (opcode >> 4) & 0x000f;
    unsigned char n = opcode & 0x000f;
    unsigned char kk = opcode & 0x00ff;

    //Look at first 4 bits
    switch (opcode & 0xf000)
    {
        //JP: Jump to location nnn
        case 0x1000:
            chip8->registers.PC = nnn;
        break;

        //CALL: Call subroutine at nnn
        case 0x2000:
            chip8_stack_push(chip8, chip8->registers.PC);
            chip8->registers.PC = nnn;
        break;

        //SE: Skip next instruction if Vx == kk
        case 0x3000:
            if (chip8->registers.V[x] == kk)
                chip8->registers.PC += 2;
        break;

        //SNE: Skip next instruction if Vx == kk
        case 0x4000:
            if (chip8->registers.V[x] != kk)
                chip8->registers.PC += 2;
        break;

        //SE: Skips next instruction if vx == vy
        case 0x5000:
            if (chip8->registers.V[x] == chip8->registers.V[y])
                chip8->registers.PC += 2;
        break;

        //LD: put value kk into Vx
        case 0x6000:
            chip8->registers.V[x] = kk;
        break;

        //ADD: adds value kk to Vx then stores it in Vx
        case 0x7000:
            chip8->registers.V[x] += kk;
        break;

        //Run boolean operations
        case 0x8000:
            chip8_exec_bool(chip8, opcode);
        break;

        //SNE: Skips next instruction if Vx != Vy
        case 0x9000:
            if (chip8->registers.V[x] != chip8->registers.V[y])
                chip8->registers.PC += 2;
        break;

        //LD: set I to nnn
        case 0xA000:
            chip8->registers.I = nnn;
        break;

        //JP: Jump to location nnn + v0
        case 0xB000:
            chip8->registers.PC = nnn+chip8->registers.V[0];
        break;

        //RND: Vx = random byte & kk
        case 0xC000:
            srand(clock());
            chip8->registers.V[x] = (rand()%255) & kk;
        break;

        //DRW: draw n byte sprite at memory location I at (Vx, Vy)
        case 0xD000:
            {
                const char* sprite = (const char*) &chip8->memory.memory[chip8->registers.I];
                chip8->registers.V[0x0f] = chip8_screen_draw_sprite(&chip8->screen,chip8->registers.V[x],chip8->registers.V[y],sprite,n);
            }
        break;

        //Keyboard
        case 0xE000:
            switch (opcode & 0x00ff)
            {
                //SKP: if key Vx is down then increment PC by 2
                case 0x9e:
                    if (chip8_keyboard_is_down(&chip8->keyboard,chip8->registers.V[x]))
                        chip8->registers.V[x] += 2;
                break;

                //SKNP if key Vx is not pressed then increment PC by 2
                case 0xA1:
                    if (!(chip8_keyboard_is_down(&chip8->keyboard,chip8->registers.V[x])))
                            chip8->registers.V[x] += 2;
                break;
            }
        break;

        case 0xF000:
            chip8_exec_f(chip8, opcode);
        break;

    }
}

void chip8_exec_bool(struct chip8* chip8, unsigned short opcode)
{   
    unsigned char x = (opcode >> 8) & 0x000f;
    unsigned char y = (opcode >> 4) & 0x000f;

    unsigned short tmp = 0;

    switch(opcode & 0x000f)
    {
        //LD: Stores value in Vy in Vx
        case 0x0000:
            chip8->registers.V[x] = chip8->registers.V[y];
        break;

        //OR: Vx = Vx || Vy
        case 0x0001:
            chip8->registers.V[x] = chip8->registers.V[y] || chip8->registers.V[x];
        break;

        //AND: Vx = Vx && Vy
        case 0x0002:
            chip8->registers.V[x] = chip8->registers.V[y] && chip8->registers.V[x];
        break;

        //XOR: Vx = Vx ^ Vy
        case 0x0003:
            chip8->registers.V[x] = chip8->registers.V[y] ^ chip8->registers.V[x];
        break;

        //ADD: Vx = Vx + Vy
        case 0x0004:
            tmp = chip8->registers.V[x] + chip8->registers.V[y];
            chip8->registers.V[0x0f] = tmp > 0xff;
            chip8->registers.V[x] = tmp;
        break;

        //SUB: Vx = Vx - Vy
        case 0x0005:
            chip8->registers.V[0x0f] = chip8->registers.V[x] > chip8->registers.V[y];
            chip8->registers.V[x] -= chip8->registers.V[y];
        break;

        //SHR: Vx = Vx SHR 1
        case 0x0006:
            chip8->registers.V[0x0f] = chip8->registers.V[x] & 0x01;
            chip8->registers.V[x] /= 2;
        break;

        //SUBN: Vx = Vy - Vx
        case 0x0007:
            chip8->registers.V[0x0f] = chip8->registers.V[y] > chip8->registers.V[x];
            chip8->registers.V[x] = chip8->registers.V[y] - chip8->registers.V[x];
        break;

        //SHL: Vx SHL 1
        case 0x000E:
            chip8->registers.V[0x0f] = chip8->registers.V[x] & 0b10000000;
            chip8->registers.V[x] *= 2;
        break;
    }   
}

void chip8_exec_f(struct chip8* chip8, unsigned short opcode)
{
    unsigned char x = (opcode >> 8) & 0x000f;
    switch (opcode & 0x00ff)
    {
        //LD: Vx = delay timer value
        case 0x07:
            chip8->registers.V[x] = chip8->registers.delay_timer;
        break;

        //LD: Wait for key press then store the value of the key in Vx
        case 0x0A:
            chip8->registers.V[x] = chip8_wait_for_key_press(chip8);
        break;

        //LD: Set delay timer
        case 0x15:
            chip8->registers.delay_timer = chip8->registers.V[x];
        break;

        //LD: Set sound timer
        case 0x18:
            chip8->registers.sound_timer = chip8->registers.V[x];
        break;

        //ADD: set I = I + Vx
        case 0x1E:
            chip8->registers.I += chip8->registers.V[x];
        break;

        //LD: set I = location of sprite for digit Vx
        case 0x29:
            chip8->registers.I = chip8->registers.V[x]*CHIP8_DEFAULT_SPRITE_HEIGHT;
        break;

        //LD: store BCD representation of Vx in memory locations I, I+1 and I+2
        case 0x33:
            {
                unsigned char hundreds = chip8->registers.V[x]/100;
                unsigned char tens = chip8->registers.V[x]/10 % 10;
                unsigned char units = chip8->registers.V[x] % 10;

                chip8_memory_set(&chip8->memory,chip8->registers.I, hundreds);
                chip8_memory_set(&chip8->memory,chip8->registers.I+1, tens);
                chip8_memory_set(&chip8->memory,chip8->registers.I+2, units);
            }
        break;

        //LD: Store registers v0 through vx in memory starting location I
        case 0x55:
            {
                for (int i=0;i<=x;i++)
                {
                    chip8_memory_set(&chip8->memory,chip8->registers.I+i, chip8->registers.V[i]);
                }               
            }

        break;

        //LD: Read registers v0 through vx from memory starting location I
        case 0x65:
            {
                for (int i=0;i<=x;i++)
                {
                    chip8->registers.V[i] = chip8_memory_get(&chip8->memory,chip8->registers.I+i);
                }
            }
        break;
    }
}

char chip8_wait_for_key_press(struct chip8* chip8)
{
    SDL_Event event;
    while (SDL_WaitEvent(&event))
    {
        if (event.type != SDL_KEYDOWN) continue;

        char c = event.key.keysym.sym;
        char chip8_key = chip8_keyboard_map(&chip8->keyboard,c);
        if (chip8_key != -1) return chip8_key;
    }

    return -1;
}