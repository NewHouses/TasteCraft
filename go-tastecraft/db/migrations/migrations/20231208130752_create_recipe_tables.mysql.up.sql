CREATE TABLE dish_types (
  dish_type_id INT NOT NULL AUTO_INCREMENT,
  dish_type VARCHAR(45) NOT NULL,
  PRIMARY KEY (dish_type_id)
);

CREATE TABLE recipes (
  recipe_id INT NOT NULL AUTO_INCREMENT,
  recipe_name VARCHAR(45) NOT NULL,
  dish_type INT NOT NULL DEFAULT 1,
  PRIMARY KEY (recipe_id),
  UNIQUE (recipe_name),
  INDEX dish_type_recipe_id_idx (dish_type ASC) VISIBLE,
  CONSTRAINT dish_type_recipe_id
    FOREIGN KEY (dish_type)
    REFERENCES dish_types (dish_type_id)
    ON DELETE NO ACTION
    ON UPDATE CASCADE);


CREATE TABLE foods (
  food_id INT NOT NULL AUTO_INCREMENT,
  food_name VARCHAR(45) NOT NULL DEFAULT "",
  UNIQUE (food_name),
  PRIMARY KEY (food_id)
);

CREATE TABLE ingredients (
  id INT NOT NULL AUTO_INCREMENT,
  recipe_id INT NOT NULL,
  food_id INT NOT NULL,
  quantity INT UNSIGNED NOT NULL DEFAULT 0,
  measurement VARCHAR(45) NOT NULL DEFAULT "grams",
  PRIMARY KEY (id),
  INDEX recipe_id_idx (recipe_id ASC) VISIBLE,
  CONSTRAINT recipe_ingredient_id
    FOREIGN KEY (recipe_id)
    REFERENCES recipes (recipe_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT food_ingredient_id
    FOREIGN KEY (food_id)
    REFERENCES foods (food_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE nutrional_values (
  id INT NOT NULL AUTO_INCREMENT,
  food_id INT NOT NULL,
  kilocalories INT UNSIGNED NOT NULL DEFAULT 0,
  proteins DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  carbohydrates DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  fat DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  fiber DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  calcium DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  iron DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  zinc DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_a DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_b DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_b1 DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_b2 DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_b3 DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_b6 DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_b12 DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_c DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_d DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_e DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  vitamin_k DECIMAL UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  INDEX food_id_idx (food_id ASC) VISIBLE,
  CONSTRAINT food_nutrional_value_id
    FOREIGN KEY (food_id)
    REFERENCES foods (food_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

CREATE TABLE steps (
  id INT NOT NULL AUTO_INCREMENT,
  recipe_id INT NOT NULL,
  step VARCHAR(120) NOT NULL DEFAULT "",
  PRIMARY KEY (id),
  INDEX recipe_step_id_idx (recipe_id ASC) VISIBLE,
  CONSTRAINT recipe_step_id
    FOREIGN KEY (recipe_id)
    REFERENCES recipes (recipe_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

INSERT INTO dish_types (dish_type) VALUES ('generic');
INSERT INTO dish_types (dish_type) VALUES ('fish');
INSERT INTO dish_types (dish_type) VALUES ('seafood');
INSERT INTO dish_types (dish_type) VALUES ('meat');
INSERT INTO dish_types (dish_type) VALUES ('vegetable');
INSERT INTO dish_types (dish_type) VALUES ('pasta');
INSERT INTO dish_types (dish_type) VALUES ('cereal');
INSERT INTO dish_types (dish_type) VALUES ('legume');
INSERT INTO dish_types (dish_type) VALUES ('soup');
INSERT INTO dish_types (dish_type) VALUES ('snack');
INSERT INTO dish_types (dish_type) VALUES ('dessert');