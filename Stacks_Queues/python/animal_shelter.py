from datetime import datetime
from exceptions import NoAnimalsException, NoDogsException, NoCatsException


class Animal(object):
    def __init__(self, name):
        self.name = name
        self.arrive_time = datetime.now()


class Dog(Animal):
    def __str__(self):
        return f"Dog {self.name} arrived at {self.arrive_time}"


class Cat(Animal):
    def __str__(self):
        return f"Cat {self.name} arrived at {self.arrive_time}"


class AnimalShelter(object):
    def __init__(self):
        self.dogs = []
        self.cats = []

    def enque(self, animal: Animal):
        if type(animal) == Dog:
            self.dogs.append(animal)
        elif type(animal) == Cat:
            self.cats.append(animal)


    def deque_any(self):
        if not self.dogs:
            if not self.cats:
                raise NoAnimalsException()

            first = self.cats[0]
            self.cats.remove(first)

        elif not self.cats:
            first = self.dogs[0]
            self.dogs.remove(first)

        else:
            first_dog = self.dogs[0]
            first_cat = self.cats[0]

            if first_dog.arrive_time <= first_cat.arrive_time:
                first = self.dogs[0]
                self.dogs.remove(first)
            else:
                first = self.cats[0]
                self.cats.remove(first)

        return first


    def deque_dog(self):
        if not self.dogs:
            raise NoDogsException()

        first = self.dogs[0]
        self.dogs.remove(first)
        return first


    def deque_cat(self):
        if not self.cats:
            raise NoCatsException()

        first = self.cats[0]
        self.cats.remove(first)
        return first

    def __str__(self):
        return f"\n".join([f"{animal.name}: {animal.arrive_time}" for animal in self.dogs + self.cats])



def main():
    shelter = AnimalShelter()

    dog = Dog("Barbos")
    cat = Cat("Murzik")
    shelter.enque(dog)
    shelter.enque(cat)
    print(shelter)
    print(shelter.deque_any())
    print(shelter)


if __name__ == '__main__':
    main()
